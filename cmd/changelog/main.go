/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
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

// getLatestTag retrieves the most recent Git tag in the current repository.
// It returns the latest tag as a string.
func getLatestTag() (string, error) {

	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}

// getCommitsSinceLastTag retrieves the list of commit messages since the specified Git tag up to HEAD.
// Takes a tag string as input and returns a slice of commit messages or an error if retrieval fails.
func getCommitsSinceLastTag(tag string) ([]string, error) {

	cmd := exec.Command("git", "log", fmt.Sprintf("%s..HEAD", tag),
		"--oneline",
		"--no-decorate",
		"--format=%s")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	commits := strings.Split(strings.TrimSpace(out.String()), "\n")
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

// generateChangelog processes a list of commit messages and organizes them into categorized changelog sections.
func generateChangelog(tag string, commits []string) string {

	var changelog strings.Builder

	sections := map[string]*changelogSection{}

	for _, commit := range commits {

		matches := reConventionalCommit.FindStringSubmatch(commit)
		if matches == nil {
			continue
		}

		commitType := matches[1]
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

// nextMinorVersion calculates the next minor version based on a given semantic version tag.
// The function expects a version tag in the format 'vMAJOR.MINOR.PATCH' and returns the incremented minor version.
func nextMinorVersion(tag string, hotfix bool) (string, error) {

	tag = strings.TrimPrefix(tag, "v")

	parts := strings.Split(tag, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version tag format")
	}

	// parse MAJOR, MINOR, and PATCH
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}

	if hotfix {
		patch += 1
	} else {
		minor += 1
		patch = 0
	}

	return fmt.Sprintf("v%d.%d.%d", major, minor, patch), nil
}

func main() {
	hotfix := flag.Bool("hotfix", false, "Calculate the next PATCH version instead of MINOR")
	flag.Parse()

	latestTag, err := getLatestTag()
	if err != nil {
		fmt.Println("Error fetching latest tag:", err)
		os.Exit(1)
	}

	nextTag, err := nextMinorVersion(latestTag, *hotfix)
	if err != nil {
		fmt.Println("Error generating next minor version:", err)
	}

	// Get commits since the last tag
	commits, err := getCommitsSinceLastTag(latestTag)
	if err != nil {
		fmt.Println("Error fetching commits:", err)
		os.Exit(1)
	}

	// Generate changelog entry
	changelogEntry := generateChangelog(nextTag, commits)

	// Print or save the changelog
	fmt.Print("\n", changelogEntry)
}
