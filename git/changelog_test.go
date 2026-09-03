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

func TestRenderChangelogTypeMapping(t *testing.T) {
	t.Run("CustomTypeIntoExistingSection", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"cicd(ci): bump runner"}, nil, nil,
			WithTypeMapping(map[string]string{"cicd": "Changed"}))
		xt.Assert(t, strings.Contains(got, "### Changed\n\n- **ci**: bump runner"),
			"expected cicd rendered under Changed, got %q", got)
	})

	t.Run("CustomTypeUnmappedIsDropped", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"cicd(ci): bump runner"}, nil, nil)
		xt.Assert(t, !strings.Contains(got, "bump runner"), "expected unmapped type dropped, got %q", got)
	})

	t.Run("CustomTypeMultiScope", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"cicd(ci|build): bump runner"}, nil, nil,
			WithTypeMapping(map[string]string{"cicd": "Changed"}))
		xt.Assert(t, strings.Contains(got, "- **ci**: bump runner"),
			"expected first scope kept, got %q", got)
	})

	t.Run("BuiltInTypeOverridden", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"docs: explain it"}, nil, nil,
			WithTypeMapping(map[string]string{"docs": "Documentation"}))
		xt.Assert(t, strings.Contains(got, "### Documentation\n\n- explain it"),
			"expected docs remapped, got %q", got)
	})

	t.Run("EmptyEntriesIgnored", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"cicd: bump runner"}, nil, nil,
			WithTypeMapping(map[string]string{"": "Changed", "cicd": ""}))
		xt.Assert(t, !strings.Contains(got, "bump runner"), "expected empty entries ignored, got %q", got)
	})

	t.Run("DefaultsNotMutated", func(t *testing.T) {
		_ = RenderChangelog("v1.2.0", []string{"cicd: bump runner"}, nil, nil,
			WithTypeMapping(map[string]string{"cicd": "Changed"}))

		got := RenderChangelog("v1.2.0", []string{"cicd: bump runner"}, nil, nil)
		xt.Assert(t, !strings.Contains(got, "bump runner"),
			"expected built-in mapping unchanged, got %q", got)
	})

	t.Run("SkipTypesAppliesToCustomType", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", []string{"cicd: bump runner"}, []string{"cicd"}, nil,
			WithTypeMapping(map[string]string{"cicd": "Changed"}))
		xt.Assert(t, !strings.Contains(got, "bump runner"), "expected cicd skipped, got %q", got)
	})
}

func TestRenderChangelogScopeOrder(t *testing.T) {
	t.Run("UnscopedFirstThenSorted", func(t *testing.T) {
		commits := []string{
			"feat(sdk): add sdk thing",
			"feat(api): add api thing",
			"feat: add plain thing",
			"feat(cli): add cli thing",
		}

		got := RenderChangelog("v1.2.0", commits, nil, nil)
		want := "### Added\n\n" +
			"- add plain thing\n" +
			"- **api**: add api thing\n" +
			"- **cli**: add cli thing\n" +
			"- **sdk**: add sdk thing\n"

		xt.Assert(t, strings.Contains(got, want), "expected %q, got %q", want, got)
	})

	t.Run("Deterministic", func(t *testing.T) {
		commits := []string{
			"feat(sdk): add sdk thing",
			"feat: add plain thing",
			"feat(api): add api thing",
			"fix(cli): correct cli thing",
			"fix(api): correct api thing",
		}

		want := RenderChangelog("v1.2.0", commits, nil, nil)
		for range 20 {
			xt.Eq(t, want, RenderChangelog("v1.2.0", commits, nil, nil))
		}
	})
}

func TestRenderChangelogSectionOrder(t *testing.T) {
	commits := []string{"feat: add thing", "fix: correct thing", "deps: bump x"}

	t.Run("NewSectionAppended", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", commits, nil, nil,
			WithTypeMapping(map[string]string{"deps": "Dependencies"}))

		fixed := strings.Index(got, "### Fixed")
		deps := strings.Index(got, "### Dependencies")

		xt.Assert(t, deps > 0, "expected Dependencies section, got %q", got)
		xt.Assert(t, deps > fixed, "expected Dependencies after Fixed, got %q", got)
	})

	t.Run("Pinned", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", commits, nil, nil,
			WithTypeMapping(map[string]string{"deps": "Dependencies"}),
			WithSectionOrder("Dependencies", "Added", "Changed", "Fixed"))

		deps := strings.Index(got, "### Dependencies")
		added := strings.Index(got, "### Added")

		xt.Assert(t, deps > 0 && added > 0, "expected both sections, got %q", got)
		xt.Assert(t, deps < added, "expected Dependencies before Added, got %q", got)
	})

	t.Run("PinnedPartialAppendsRest", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", commits, nil, nil,
			WithTypeMapping(map[string]string{"deps": "Dependencies"}),
			WithSectionOrder("Fixed"))

		fixed := strings.Index(got, "### Fixed")
		added := strings.Index(got, "### Added")
		deps := strings.Index(got, "### Dependencies")

		xt.Assert(t, fixed >= 0 && added > fixed && deps > added,
			"expected Fixed, then Added and Dependencies sorted, got %q", got)
	})

	t.Run("EmptyOrderKeepsDefault", func(t *testing.T) {
		got := RenderChangelog("v1.2.0", commits, nil, nil, WithSectionOrder())

		added := strings.Index(got, "### Added")
		fixed := strings.Index(got, "### Fixed")

		xt.Assert(t, added >= 0 && fixed > added, "expected default order, got %q", got)
	})
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
