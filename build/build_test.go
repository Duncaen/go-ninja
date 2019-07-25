package build

import (
	"bytes"
	"testing"
)

var buildTests = []struct {
	in  Build
	out string
}{
	{
		Build{Rule: "foo"},
		`build : foo
`},
	{
		Build{Rule: "foo", Outputs: []string{"bar"}},
		`build bar: foo
`},
}

func TestBuild(t *testing.T) {
	for _, tt := range buildTests {
		var b bytes.Buffer
		if _, err := tt.in.WriteTo(&b); err != nil {
			t.Error(err)
			continue
		}
		res := b.String()
		if res != tt.out {
			t.Errorf("expected %q got %q", tt.out, res)
		}
	}
}
