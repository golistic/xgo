/*
 * Copyright (c) 2025, Geert JM Vanderkelen
 */

package git

import (
	"strings"
	"testing"
	"time"

	"github.com/golistic/xgo/xt"
)

func TestFirstScope(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "SingleScope",
			input: "(api)",
			want:  "api",
		},
		{
			name:  "PipeSeparated",
			input: "(api|cli)",
			want:  "api",
		},
		{
			name:  "CommaSeparated",
			input: "(api,cli)",
			want:  "api",
		},
		{
			name:  "SemicolonSeparated",
			input: "(api;cli)",
			want:  "api",
		},
		{
			name:  "SpacedSeparators",
			input: "(api | cli ; sdk)",
			want:  "api",
		},
		{
			name:  "NoScope",
			input: "",
			want:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xt.Eq(t, tc.want, firstScope(tc.input))
		})
	}
}

func TestRenderChangelog(t *testing.T) {
	heading := "## [1.2.0] - " + time.Now().Format(time.DateOnly)

	tests := []struct {
		name    string
		commits []string
		want    string
	}{
		{
			name:    "SingleScope",
			commits: []string{"feat(api): add endpoint"},
			want:    "### Added\n\n- **api**: add endpoint\n",
		},
		{
			name:    "MultipleScopesKeepFirst",
			commits: []string{"feat(api|cli): add endpoint"},
			want:    "### Added\n\n- **api**: add endpoint\n",
		},
		{
			name:    "MultipleScopesComma",
			commits: []string{"fix(cli, api): correct exit code"},
			want:    "### Fixed\n\n- **cli**: correct exit code\n",
		},
		{
			name:    "MultipleScopesSemicolon",
			commits: []string{"refactor(sdk;api;cli): move helpers"},
			want:    "### Changed\n\n- **sdk**: move helpers\n",
		},
		{
			name:    "WithoutScope",
			commits: []string{"feat: add thing"},
			want:    "### Added\n\n- add thing\n",
		},
		{
			name:    "GroupedByFirstScope",
			commits: []string{"feat(api|cli): add endpoint", "feat(api,sdk): add filter"},
			want:    "### Added\n\n- **api**:\n    - add filter\n    - add endpoint\n",
		},
		{
			name:    "NonConventionalIgnored",
			commits: []string{"whatever happened here"},
			want:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RenderChangelog("v1.2.0", tc.commits, nil, nil)
			xt.Assert(t, strings.HasPrefix(got, heading), "expected version heading")
			xt.Assert(t, strings.Contains(got, tc.want),
				"expected changelog to contain %q, got %q", tc.want, got)
		})
	}
}

func TestRenderChangelogSkipping(t *testing.T) {
	t.Run("SkipTypes", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"feat(api|cli): add endpoint"}, []string{"feat"}, nil)
		xt.Assert(t, !strings.Contains(got, "add endpoint"), "expected feat to be skipped")
	})

	t.Run("SkipFirstScope", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"feat(api|cli): add endpoint"}, nil, []string{"api"})
		xt.Assert(t, !strings.Contains(got, "add endpoint"), "expected scope api to be skipped")
	})

	t.Run("SkipSecondaryScopeHasNoEffect", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"feat(api|cli): add endpoint"}, nil, []string{"cli"})
		xt.Assert(t, strings.Contains(got, "- **api**: add endpoint"), "expected entry kept, got %q", got)
	})
}
