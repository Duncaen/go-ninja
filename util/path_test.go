package util

import (
	"testing"
)

var testPrefixPaths = []struct {
	prefix string
	paths  []string
	out    []string
}{
	{
		prefix: "foo",
		paths:  []string{"bar"},
		out:    []string{"foo/bar"},
	},
	{
		prefix: "fizz/buzz",
		paths:  []string{"foo/bar"},
		out:    []string{"fizz/buzz/foo/bar"},
	},
	{
		prefix: "fizz/buzz",
		paths:  []string{"foo", "bar"},
		out:    []string{"fizz/buzz/foo", "fizz/buzz/bar"},
	},
}

func testEQPaths(t *testing.T, got, exp []string) bool {
	t.Helper()
	if len(got) != len(exp) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if got[i] != exp[i] {
			return false
		}
	}
	return true
}

func TestPrefixPaths(t *testing.T) {
	for _, tt := range testPrefixPaths {
		ret := PrefixPaths(tt.prefix, tt.paths...)
		if !testEQPaths(t, ret, tt.out) {
			t.Errorf("expected %v, got %v", tt.out, ret)
		}
	}
}
