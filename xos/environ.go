// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"os"
)

// EnvironMap goes through all environment variables, saving them
// into a map with name being the key.
func EnvironMap() map[string]string {
	return environMapLookup(nil)
}

// EnvironMapLookup goes through all environment variables, saving them
// into a map with name being the key. When provided, the mapper function
// can be used to include variables.
//
// For example, you would like to include some Go environment variables:
//
//	envMap := EnvironMapLookup(func(name, _ string) bool {
//	  return name == "GOROOT" || name == "GOPATH"
//	})
//
// If no mapper is provided, EnvironMapLookup works exactly as EnvironMap.
func EnvironMapLookup(mapper func(name, value string) bool) map[string]string {
	return environMapLookup(mapper)
}

func environMapLookup(mapper func(name, value string) bool) map[string]string {
	m := map[string]string{}

	for _, envVar := range os.Environ() {
		for i := 0; i < len(envVar); i++ {
			if envVar[i] == '=' {
				n := envVar[0:i]
				v := envVar[i+1:]
				if mapper == nil || mapper(n, v) {
					m[n] = v
				}
			}
		}
	}

	return m
}
