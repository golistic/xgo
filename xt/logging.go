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

func (l *LogAgg) Write(p []byte) (n int, err error) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.entries = append(l.entries, string(p))

	return len(p), nil
}

func (l *LogAgg) Reset() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.entries = nil
}

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
