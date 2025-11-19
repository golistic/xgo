/*
 * Copyright (c) 2024, 2025, Geert JM Vanderkelen
 */

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/golistic/xgo/git"
)

func main() {
	var hotfix bool
	flag.BoolVar(&hotfix, "hotfix", false, "Calculate the next PATCH version instead of MINOR")

	var tagOnly bool
	flag.BoolVar(&tagOnly, "tag-only", false, "Only show the tag/version instead of the full changelog")

	var flagSkipTypes string
	flag.StringVar(&flagSkipTypes, "skip-types", "", "Comma-separated list of types (feat, fix, etc.) to skip")

	var flagSkipScopes string
	flag.StringVar(&flagSkipScopes, "skip-scopes", "", "Comma-separated list of scopes to skip")

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

	if tagOnly {
		tag, err := git.NextTag(tagBranch, hotfix)
		if err != nil {
			fmt.Println("Error calculating next tag:", err)
			return
		}
		if tag != "" {
			println(tag)
		}
		return
	}

	if out, err := git.GenerateChangelog(tagBranch, hotfix, skipTypes, skipScopes); err != nil {
		fmt.Println("Error generating changelog:", err)
	} else if out != "" {
		fmt.Print(out)
	} else {
		fmt.Println("No changes detected.")
	}
}
