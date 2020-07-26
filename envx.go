// Copyright (c) 2010 The Go Authors
// Copyright (c) 2020 Bojan Zivanovic
// SPDX-License-Identifier: BSD-2-Clause

// Package envx allows expanding env variables with defaults: ${var:default}.
package envx

import (
	"os"
	"strings"
)

// Expand replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables are replaced by the empty string. Defaults are supported
// using the ${var:default} format.
func Expand(s string) string {
	return ExpandFunc(s, os.Getenv)
}

// ExpandFunc replaces ${var} or $var in the string based on the mapping function.
// For example, envx.Expand(s) is equivalent to envx.ExpandFunc(s, os.Getenv).
// Defaults are supported using the ${var:default} format.
func ExpandFunc(s string, mapping func(string) string) string {
	var buf []byte
	// ${} is all ASCII, so bytes are fine for this operation.
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == '$' && j+1 < len(s) {
			if buf == nil {
				buf = make([]byte, 0, 2*len(s))
			}
			buf = append(buf, s[i:j]...)
			name, w := getShellName(s[j+1:])
			if name == "" && w > 0 {
				// Encountered invalid syntax; eat the
				// characters.
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else {
				// Replace the name, using a default value
				// if defined. The next 6 lines is where this
				// package diverges from stdlib.
				nameParts := strings.SplitN(name, ":", 2)
				name := nameParts[0]
				value := mapping(name)
				if value == "" && len(nameParts) == 2 {
					value = nameParts[1]
				}
				buf = append(buf, value...)
			}
			j += w
			i = j + 1
		}
	}
	if buf == nil {
		return s
	}
	return string(buf) + s[i:]
}

// getShellName returns the name that begins the string and the number of bytes
// consumed to extract it. If the name is enclosed in {}, it's part of a ${}
// expansion and two more bytes are needed than the length of the name.
func getShellName(s string) (string, int) {
	switch {
	case s[0] == '{':
		if len(s) > 2 && isShellSpecialVar(s[1]) && s[2] == '}' {
			return s[1:2], 3
		}
		// Scan to closing brace
		for i := 1; i < len(s); i++ {
			if s[i] == '}' {
				if i == 1 {
					return "", 2 // Bad syntax; eat "${}"
				}
				return s[1:i], i + 1
			}
		}
		return "", 1 // Bad syntax; eat "${"
	case isShellSpecialVar(s[0]):
		return s[0:1], 1
	}
	// Scan alphanumerics.
	var i int
	for i = 0; i < len(s) && isAlphaNum(s[i]); i++ {
	}
	return s[:i], i
}

// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}
