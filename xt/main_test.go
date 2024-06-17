/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xt

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv(EnvNoColors, "")

	os.Exit(m.Run())
}
