// Package rule implements ninja build file rule blocks
package rule

import (
	"io"
	"strings"

	"github.com/Duncaen/go-ninja"
	"github.com/Duncaen/go-ninja/util"
)

// Command represents the builds command variable
type Command []string

// Deps represents the deps variable in rules
//
// https://ninja-build.org/manual.html#ref_headers
type Deps int

const (
	// DepsGCC instructs ninja to use gcc dependency files
	DepsGCC Deps = 1 << iota
	// DepsMSVC instruct ninja to use MSVC dependency files
	DepsMSVC
)

func (d Deps) String() string {
	switch d {
	case DepsGCC: return "gcc"
	case DepsMSVC: return "msvc"
	}
	panic("unknown deps")
}

// Rule represents a ninja rule block
type Rule struct {
	Name string
	Command string
	Deps Deps
	MSVCDepsPrefix string
	Vars ninja.Vars
}

// New creates a new ninja rule block
func New(name string, v ...interface{}) Rule {
	r := Rule{Name: name}
	for _, n := range v {
		switch x := n.(type) {
		case Command:
			r.Command = strings.Join(x, " && ")
		case ninja.Var:
			r.Vars = append(r.Vars, x)
		case Deps:
			r.Deps = x
		default:
			panic("unsupported type")
		}
	}
	return r
}

// RequiredVersion returns the ninja version required for the rule
func (r Rule) RequiredVersion() (v string) {
	// deps since Ninja 1.3
	if r.Deps != 0 {
		v = "1.3"
	}
	// msvc_deps_prefix since Ninja 1.5
	if r.MSVCDepsPrefix != "" {
		v = "1.5"
	}
	return
}

func writeVar(w io.Writer, key, val string) {
	io.WriteString(w, key)
	io.WriteString(w, " = ")
	io.WriteString(w, val)
	io.WriteString(w, "\n")
}

// WriteTo writes the rule block to w
func (r Rule) WriteTo(w io.Writer) (n int64, err error) {
	wr := util.NewSeqWriter(w)
	io.WriteString(wr, "rule ")
	io.WriteString(wr, r.Name)
	io.WriteString(wr, "\n")
	w2 := util.NewPrefixWriter(wr, []byte("  "))
	writeVar(w2, "command", r.Command)
	if r.Deps != 0 {
		writeVar(w2, "deps", r.Deps.String())
	}
	r.Vars.WriteTo(w2)
	return wr.Done()
}
