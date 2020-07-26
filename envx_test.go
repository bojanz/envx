// Copyright (c) 2010 The Go Authors
// Copyright (c) 2020 Bojan Zivanovic
// SPDX-License-Identifier: BSD-2-Clause

package envx_test

import (
	"os"
	"testing"

	"github.com/bojanz/envx"
)

// testGetenv gives us a controlled set of variables for testing ExpandFunc.
func testGetenv(s string) string {
	switch s {
	case "*":
		return "all the args"
	case "#":
		return "NARGS"
	case "$":
		return "PID"
	case "1":
		return "ARGUMENT1"
	case "HOME":
		return "/usr/gopher"
	case "H":
		return "(Value of H)"
	case "home_1":
		return "/usr/foo"
	case "_":
		return "underscore"
	}
	return ""
}

var expandTests = []struct {
	in, out string
}{
	// stdlib tests
	{"", ""},
	{"$*", "all the args"},
	{"$$", "PID"},
	{"${*}", "all the args"},
	{"$1", "ARGUMENT1"},
	{"${1}", "ARGUMENT1"},
	{"now is the time", "now is the time"},
	{"$HOME", "/usr/gopher"},
	{"$home_1", "/usr/foo"},
	{"${HOME}", "/usr/gopher"},
	{"${H}OME", "(Value of H)OME"},
	{"A$$$#$1$H$home_1*B", "APIDNARGSARGUMENT1(Value of H)/usr/foo*B"},
	{"start$+middle$^end$", "start$+middle$^end$"},
	{"mixed$|bag$$$", "mixed$|bagPID$"},
	{"$", "$"},
	{"$}", "$}"},
	{"${", ""},  // invalid syntax; eat up the characters
	{"${}", ""}, // invalid syntax; eat up the characters

	// Defaults
	{"${HOME:/srv}", "/usr/gopher"},
	{"${WORKDIR:/srv}", "/srv"},
	{"${APP_NAME}", ""},
	{"${APP_NAME:My App}", "My App"},
	{"${PORT:80}", "80"},
	{"${ADDR::80}", ":80"},
}

func TestGet(t *testing.T) {
	v := envx.Get("LISTEN", "0.0.0.0:80")
	if v != "0.0.0.0:80" {
		t.Errorf("Get(LISTEN)=%q; expected %q", v, "0.0.0.0:80")
	}

	os.Setenv("LISTEN", "0.0.0.0:443")
	v = envx.Get("LISTEN", "0.0.0.0:443")
	if v != "0.0.0.0:443" {
		t.Errorf("Get(LISTEN)=%q; expected %q", v, "0.0.0.0:443")
	}
	os.Unsetenv("LISTEN")
}

func TestExpand(t *testing.T) {
	for _, test := range expandTests {
		result := envx.ExpandFunc(test.in, testGetenv)
		if result != test.out {
			t.Errorf("Expand(%q)=%q; expected %q", test.in, result, test.out)
		}
	}
}
