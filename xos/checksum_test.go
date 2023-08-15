// Copyright (c) 2023, Geert JM Vanderkelen

package xos_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/golistic/xgo/xt"

	"github.com/golistic/xgo/xos"
)

func TestFChecksumMD5(t *testing.T) {
	fp, err := os.Open(filepath.Join("_testdata", "gopher_1up.png"))
	xt.OK(t, err)
	defer func() { _ = fp.Close() }()

	exp := "6620559b36ac9ed6765fc885d045593c"
	have, err := xos.FChecksumMD5(fp)

	xt.Eq(t, exp, have)

	havePos, err := fp.Seek(0, io.SeekCurrent)
	xt.OK(t, err)
	xt.Eq(t, 0, havePos)
}

func TestFChecksumSHA1(t *testing.T) {
	fp, err := os.Open(filepath.Join("_testdata", "gopher_1up.png"))
	xt.OK(t, err)
	defer func() { _ = fp.Close() }()

	exp := "b1b437d0ee444ad7585c8f5cc82e87eafed020a5"
	have, err := xos.FChecksumSHA1(fp)

	xt.Eq(t, exp, have)

	havePos, err := fp.Seek(0, io.SeekCurrent)
	xt.OK(t, err)
	xt.Eq(t, 0, havePos)
}

func TestFChecksumSHA256(t *testing.T) {
	fp, err := os.Open(filepath.Join("_testdata", "gopher_1up.png"))
	xt.OK(t, err)
	defer func() { _ = fp.Close() }()

	exp := "bde6935a7072d1d5e04a1dacc433293d9506cb80e21058a3f78dc19a232c98e0"
	have, err := xos.FChecksumSHA256(fp)

	xt.Eq(t, exp, have)

	havePos, err := fp.Seek(0, io.SeekCurrent)
	xt.OK(t, err)
	xt.Eq(t, 0, havePos)
}
