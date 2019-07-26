package ninja

import (
	"testing"
)

var escapeTests = []struct {
	in  string
	out string
}{
	{"foo/ bar", "foo/$ bar"},
	{"foo/:bar", "foo/$:bar"},
	{"foo/\nbar", "foo/$\nbar"},
	{"foo/$bar", "foo/$$bar"},
}

func TestEscape(t *testing.T) {
	for _, tt := range escapeTests {
		res := Escape(tt.in)
		if res != tt.out {
			t.Errorf("got %q, expected %q", res, tt.out)
		}
	}
}
