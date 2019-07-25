package util

import (
	"testing"
	"bytes"
)

var prefixTests = []struct {
	pre string
	in []string
	out string
}{
	{">>", []string{"a = b\n"}, ">>a = b\n"},
	{">>", []string{"a = b\n", "c = d\n"}, ">>a = b\n>>c = d\n"},
	{">>", []string{"a =", " b\n"}, ">>a = b\n"},
	{">>", []string{"a", " ", "=", " ", "b", "\n"}, ">>a = b\n"},
	{">>", []string{"\n"}, ">>\n"},
	{">>", []string{"\n\n"}, ">>\n>>\n"},
	{">>", []string{"\n", "\n"}, ">>\n>>\n"},
	{">>", []string{"\n", "a"}, ">>\n>>a"},
}

func TestPrefixWriter(t *testing.T) {
	for _, tt := range prefixTests {
		var b bytes.Buffer
		wr := NewPrefixWriter(&b, []byte(tt.pre))
		for _, s := range tt.in {
			wr.Write([]byte(s))
		}
		res := b.String()
		if res != tt.out {
			t.Errorf("expected %q got %q", tt.out, res)
		}
	}
}
