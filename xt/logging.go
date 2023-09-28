/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xt

import (
	"encoding/json"
	"io"
	"regexp"
	"sync"
	"testing"
)

func NewLogAgg() *LogAgg {
	return &LogAgg{}
}

// LogAgg is a very basic log aggregation writer which can be used to find
// entries. This is really just useful for tests and also not really specific to logging.
type LogAgg struct {
	entries []string
	mu      sync.RWMutex
}

var _ io.Writer = (*LogAgg)(nil)

// Write stores entry.
func (l *LogAgg) Write(entry []byte) (n int, err error) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.entries = append(l.entries, string(entry))

	return len(entry), nil
}

// Reset clears all entries.
func (l *LogAgg) Reset() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.entries = nil
}

// Find searches for an entry which matches pattern.
func (l *LogAgg) Find(t *testing.T, pattern string) string {

	t.Helper()

	re := regexp.MustCompile(pattern)

	l.mu.Lock()
	defer l.mu.Unlock()

	var line string
	for _, line = range l.entries {
		if re.MatchString(line) {
			break
		}
	}

	return line
}

// FindJSON searches for an entry which matches pattern and returns the decoded
// JSON line as map[string]any.
func (l *LogAgg) FindJSON(t *testing.T, pattern string) map[string]any {

	t.Helper()

	line := l.Find(t, pattern)

	if line == "" {
		t.Fatalf("failed finding entry matching %s", pattern)
	}

	result := map[string]any{}
	OK(t, json.Unmarshal([]byte(line), &result))

	return result
}

// Entries returns copy of all entries.
func (l *LogAgg) Entries() []string {
	return l.entries
}
