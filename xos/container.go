// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"bytes"
	"fmt"
	"os"
)

func linuxProcControlGroup(pid int) []byte {
	data, _ := os.ReadFile(fmt.Sprintf("/proc/%d/cgroup", pid)) // error skipped; we use zero value
	return data
}

func linuxProcEnviron(pid int) []byte {
	data, _ := os.ReadFile(fmt.Sprintf("/proc/%d/environ", pid)) // error skipped; we use zero value
	return data
}

// InDockerContainer returns true if the application is running within a Docker container.
func InDockerContainer() bool {
	const pid = 1
	return IsRegularFile("/.dockerenv") ||
		bytes.Contains(linuxProcControlGroup(pid), []byte(":/docker/")) ||
		bytes.Contains(linuxProcEnviron(pid), []byte("container=lxc"))

}
