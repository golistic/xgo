// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func FChecksumMD5(fp io.ReadSeeker) (string, error) {
	const errMsg = "generating MD5 checksum (%w)"

	h := md5.New()
	_, err := io.Copy(h, fp)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}

	c := hex.EncodeToString(h.Sum(nil))

	_, err = fp.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}

	return c, nil
}

func FChecksumSHA1(fp io.ReadSeeker) (string, error) {
	const errMsg = "generating SHA1 checksum (%w)"

	h := sha1.New()
	_, err := io.Copy(h, fp)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}

	c := hex.EncodeToString(h.Sum(nil))

	_, err = fp.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}

	return c, nil
}

func FChecksumSHA256(fp io.ReadSeeker) (string, error) {
	const errMsg = "generating SHA256 checksum (%w)"

	h := sha256.New()
	_, err := io.Copy(h, fp)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}

	c := hex.EncodeToString(h.Sum(nil))

	_, err = fp.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}

	return c, nil
}
