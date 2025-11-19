/*
 * Copyright (c) 2025, Geert JM Vanderkelen
 */

package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

var reConventionalCommit = regexp.MustCompile(`^(feat|fix|hotfix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([a-zA-Z0-9_-]+\))?: (.*)$`)

var conventionalMapping = map[string]string{
	"feat":     "Added",
	"fix":      "Fixed",
	"hotfix":   "Fixed",
	"docs":     "Changed",
	"style":    "Changed",
	"refactor": "Changed",
	"perf":     "Changed",
	"build":    "Changed",
}

var sectionOrder = []string{"Added", "Changed", "Fixed"}

// LatestTag retrieves the most recent Git tag reachable from the given branch.
// If branch is empty, it uses "main" by default. It returns the latest tag as a string.
func LatestTag(branch string) (string, error) {

	if strings.TrimSpace(branch) == "" {
		branch = "main"
	}

	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0", branch)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// CommitsSince retrieves the list of commit messages since the specified Git tag up to HEAD.
func CommitsSince(tag string) ([]string, error) {

	cmd := exec.Command("git", "log", fmt.Sprintf("%s..HEAD", tag),
		"--oneline",
		"--no-decorate",
		"--format=%s")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	commits := strings.Split(strings.TrimSpace(out.String()), "\n")

	// Handle case with no commits (empty string results in [""])
	if len(commits) == 1 && commits[0] == "" {
		return []string{}, nil
	}

	return commits, nil
}

type changelogEntry struct {
	scope   string
	message string
}

type changelogSection struct {
	name    string
	entries map[string][]changelogEntry
}

func newChangelogSection(name string) *changelogSection {

	return &changelogSection{
		name:    name,
		entries: make(map[string][]changelogEntry),
	}
}

func (s *changelogSection) addEntry(name, message string) {

	if s.entries == nil {
		s.entries = make(map[string][]changelogEntry)
	}

	if slices.ContainsFunc(s.entries[name], func(e changelogEntry) bool {
		return strings.Contains(e.message, message)
	}) {
		return
	}
	s.entries[name] = append(s.entries[name], changelogEntry{
		scope:   name,
		message: message,
	})
}

// RenderChangelog processes a list of commit messages and organizes them into
// categorized changelog sections. The tag parameter is the next version tag to
// include in the heading. Use skipTypes to omit certain Conventional Commit
// types and skipScopes to omit specific scopes from the output.
func RenderChangelog(tag string, commits []string, skipTypes []string, skipScopes []string) string {

	var changelog strings.Builder

	sections := map[string]*changelogSection{}

	for _, commit := range commits {
		matches := reConventionalCommit.FindStringSubmatch(commit)
		if matches == nil {
			continue
		}

		commitType := matches[1]
		if slices.Contains(skipTypes, commitType) {
			continue
		}

		commitScope := ""
		commitMessage := ""

		switch len(matches) {
		case 4:
			commitScope = strings.Trim(matches[2], "()")
			commitMessage = matches[3]
		case 3:
			commitMessage = strings.TrimSpace(matches[2])
		default:
			continue
		}

		if header, exists := conventionalMapping[commitType]; exists {
			s, ok := sections[header]
			if !ok {
				s = newChangelogSection(header)
				sections[header] = s
			}
			s.addEntry(commitScope, strings.TrimSpace(commitMessage))
		}
	}

	changelog.WriteString(fmt.Sprintf("## [%s] - %s\n\n",
		strings.TrimPrefix(tag, "v"), time.Now().Format(time.DateOnly)))

	for _, s := range sectionOrder {
		section, ok := sections[s]
		if !ok || len(section.entries) == 0 {
			continue
		}

		changelog.WriteString(fmt.Sprintf("### %s\n\n", section.name))

		for scope, entries := range section.entries {
			if slices.Contains(skipScopes, scope) {
				continue
			}
			if len(entries) == 0 {
				continue
			}
			if scope == "" {
				for _, entry := range entries {
					changelog.WriteString(fmt.Sprintf("- %s\n", entry.message))
				}
			} else {
				if len(entries) == 1 {
					changelog.WriteString(fmt.Sprintf("- **%s**: %s\n", scope, entries[0].message))
				} else {
					slices.Reverse(entries)
					changelog.WriteString(fmt.Sprintf("- **%s**:\n", scope))
					for _, entry := range entries {
						changelog.WriteString(fmt.Sprintf("    - %s\n", entry.message))
					}
				}
			}
		}

		changelog.WriteString("\n")
	}

	return changelog.String()
}

// NextVersion calculates the next version using golang.org/x/mod/semver.
// It expects a semantic version tag like 'vMAJOR.MINOR.PATCH'.
// When hotfix is true, the PATCH component is incremented; otherwise the MINOR
// component is incremented and PATCH reset to 0. The returned version is a
// canonical semver starting with 'v'.
func NextVersion(tag string, hotfix bool) (string, error) {

	// ensure tag has the leading 'v' required by x/mod/semver
	if !strings.HasPrefix(tag, "v") {
		tag = "v" + tag
	}

	if !semver.IsValid(tag) {
		return "", fmt.Errorf("invalid semantic version: %q", tag)
	}

	// normalize to canonical form
	ver := semver.Canonical(tag)

	// strip prerelease and build metadata, operate on base version only
	if pre := semver.Prerelease(ver); pre != "" {
		ver = strings.TrimSuffix(ver, pre)
	}
	if build := semver.Build(ver); build != "" {
		ver = strings.TrimSuffix(ver, build)
	}

	// extract numeric parts
	core := strings.TrimPrefix(ver, "v")
	parts := strings.Split(core, ".")
	// be tolerant if tag is missing components (though module tags should not)
	for len(parts) < 3 {
		parts = append(parts, "0")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid major in %q: %w", tag, err)
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid minor in %q: %w", tag, err)
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", fmt.Errorf("invalid patch in %q: %w", tag, err)
	}

	if hotfix {
		patch++
	} else {
		minor++
		patch = 0
	}

	next := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	// Return canonical just in case
	return semver.Canonical(next), nil
}

// NextTag returns the next semantic version tag based on the latest tag found
// on the given branch. When hotfix is true, a PATCH increment is performed; otherwise
// MINOR is incremented and PATCH reset to 0.
func NextTag(tagBranch string, hotfix bool) (string, error) {
	latestTag, err := LatestTag(tagBranch)
	if err != nil {
		return "", fmt.Errorf("latest tag: %w", err)
	}
	nextTag, err := NextVersion(latestTag, hotfix)
	if err != nil {
		return "", fmt.Errorf("next version: %w", err)
	}
	return nextTag, nil
}

// GenerateChangelog is a high-level helper that orchestrates generating the
// changelog text based on common inputs you would pass via CLI flags.
//
// It performs the following steps:
//  1. Determine the latest tag reachable from the given branch.
//  2. Compute the next version (minor bump by default, patch when hotfix).
//  3. Collect commits since the latest tag and render the changelog.
//
// Parameters:
//   - tagBranch: branch on which to search the latest tag (defaults to "main" if empty)
//   - hotfix: whether to bump PATCH instead of MINOR
//   - skipTypes: Conventional Commit types to omit (e.g., feat, fix, chore)
//   - skipScopes: scopes to omit from the output
//
// Returns the rendered changelog, or an error.
func GenerateChangelog(tagBranch string, hotfix bool, skipTypes []string, skipScopes []string) (string, error) {

	latestTag, err := LatestTag(tagBranch)
	if err != nil {
		return "", fmt.Errorf("latest tag: %w", err)
	}

	nextTag, err := NextVersion(latestTag, hotfix)
	if err != nil {
		return "", fmt.Errorf("next version: %w", err)
	}

	commits, err := CommitsSince(latestTag)
	if err != nil {
		return "", fmt.Errorf("commits since %s: %w", latestTag, err)
	}

	entry := RenderChangelog(nextTag, commits, skipTypes, skipScopes)
	return entry, nil
}
