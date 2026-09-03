/*
 * Copyright (c) 2024, 2025, Geert JM Vanderkelen
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golistic/xgo/git"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	var hotfix bool
	flag.BoolVar(&hotfix, "hotfix", false, "Calculate the next PATCH version instead of MINOR")

	var tagOnly bool
	flag.BoolVar(&tagOnly, "tag-only", false, "Only show the tag/version instead of the full changelog")

	var flagSkipTypes string
	flag.StringVar(&flagSkipTypes, "skip-types", "", "Comma-separated list of types (feat, fix, etc.) to skip")

	var flagSkipScopes string
	flag.StringVar(&flagSkipScopes, "skip-scopes", "", "Comma-separated list of scopes to skip")

	var flagTypeMap string
	flag.StringVar(&flagTypeMap, "type-map", "",
		"Comma-separated list of extra type=section pairs (for example cicd=Changed,deps=Dependencies)")

	var flagSectionOrder string
	flag.StringVar(&flagSectionOrder, "section-order", "",
		"Comma-separated list of sections, setting the order in which they are rendered")

	var tagBranch string
	flag.StringVar(&tagBranch, "tag-branch", "main", "Branch to search for the latest tag (default: main)")

	flag.Parse()

	var skipTypes []string
	if flagSkipTypes != "" {
		skipTypes = strings.Split(flagSkipTypes, ",")
	}

	var skipScopes []string
	if flagSkipScopes != "" {
		skipScopes = strings.Split(flagSkipScopes, ",")
	}

	var opts []git.ChangelogOption

	if flagTypeMap != "" {
		typeMap, err := parseTypeMap(flagTypeMap)
		if err != nil {
			return fmt.Errorf("parsing type map: %w", err)
		}
		opts = append(opts, git.WithTypeMapping(typeMap))
	}

	if flagSectionOrder != "" {
		opts = append(opts, git.WithSectionOrder(strings.Split(flagSectionOrder, ",")...))
	}

	if tagOnly {
		tag, err := git.NextTag(tagBranch, hotfix)
		if err != nil {
			return fmt.Errorf("calculating next tag: %w", err)
		}
		if tag != "" {
			fmt.Println(tag)
		}
		return nil
	}

	out, err := git.GenerateChangelog(tagBranch, hotfix, skipTypes, skipScopes, opts...)
	if err != nil {
		return fmt.Errorf("generating changelog: %w", err)
	}

	if out == "" {
		fmt.Fprintln(os.Stderr, "No changes detected.")
		return nil
	}

	fmt.Print(out)

	return nil
}

// parseTypeMap turns a comma-separated list of type=section pairs into the
// mapping taken by git.WithTypeMapping.
func parseTypeMap(value string) (map[string]string, error) {

	mapping := map[string]string{}

	for _, pair := range strings.Split(value, ",") {
		if strings.TrimSpace(pair) == "" {
			continue
		}

		commitType, section, found := strings.Cut(pair, "=")
		commitType = strings.TrimSpace(commitType)
		section = strings.TrimSpace(section)

		if !found || commitType == "" || section == "" {
			return nil, fmt.Errorf("expected type=section, got %q", strings.TrimSpace(pair))
		}

		mapping[commitType] = section
	}

	return mapping, nil
}
