// Package build implements ninja build file build blocks
package build

import (
	"io"
	"regexp"

	"github.com/Duncaen/go-ninja"
	"github.com/Duncaen/go-ninja/util"
)

var escapeChars = regexp.MustCompile("([$: ])")

// escapePath escapes special characters in path with a $, i.e. : becomes $:
func escapePath(path string) string {
	return escapeChars.ReplaceAllString(path, "$$$1")
}

// Out represents the builds output files
type Out []string

// In represents the builds input files
type In []string

// ImplicitOut represents the builds implicit input files
type ImplicitOut []string

// ImplicitIn represents the builds implicit input files
type ImplicitIn []string

// OrderOnlyIn represents the builds order only input files
type OrderOnlyIn []string

// Build reperesents a ninja build block
type Build struct {
	Rule            string
	Outputs         []string
	ImplicitOutputs []string // since Ninja 1.7
	Inputs          []string
	ImplicitInputs  []string
	OrderOnlyInputs []string
	Vars ninja.Vars
}

// New creates a new ninja build block
func New(rule string, v ...interface{}) Build {
	b := Build{Rule: rule}
	for _, n := range v {
		switch x := n.(type) {
		case In:
			b.Inputs = append(b.Inputs, x...)
		case ImplicitIn:
			b.ImplicitInputs = append(b.ImplicitInputs, x...)
		case OrderOnlyIn:
			b.OrderOnlyInputs = append(b.OrderOnlyInputs, x...)
		case ImplicitOut:
			b.ImplicitOutputs = append(b.ImplicitOutputs, x...)
		case Out:
			b.Outputs = append(b.Outputs, x...)
		case ninja.Var:
			b.Vars = append(b.Vars, x)
		default:
			panic("unsupported type")
		}
	}
	return b
}

// writePaths writes escaped paths to w
func writePaths(w io.Writer, paths []string) {
	len := len(paths)
	for i, p := range paths {
		io.WriteString(w, escapePath(p))
		if i < len-1 {
			io.WriteString(w, " ")
		}
	}
}

// WriteTo writes the build block to w
func (b Build) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)

	io.WriteString(wr, "build ")
	writePaths(wr, b.Outputs)

	if len(b.ImplicitOutputs) > 0 {
		io.WriteString(wr, " | ")
		writePaths(wr, b.ImplicitOutputs)
	}

	io.WriteString(wr, ": ")
	io.WriteString(wr, b.Rule)

	if len(b.Inputs) > 0 {
		io.WriteString(wr, " ")
		writePaths(wr, b.Inputs)
	}

	if len(b.ImplicitInputs) > 0 {
		io.WriteString(wr, " | ")
		writePaths(wr, b.ImplicitInputs)
	}

	if len(b.OrderOnlyInputs) > 0 {
		io.WriteString(wr, " || ")
		writePaths(wr, b.OrderOnlyInputs)
	}

	io.WriteString(wr, "\n")
	b.Vars.WriteTo(util.NewPrefixWriter(wr, []byte("  ")))

	return wr.Done()
}
