package ninja

import (
	"io"

	"github.com/Duncaen/go-ninja/util"
)

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
	case DepsGCC:
		return "gcc"
	case DepsMSVC:
		return "msvc"
	}
	panic("unknown deps")
}

// Rule represents a ninja rule block
type Rule struct {
	Name           string
	Command        string // required
	Deps           Deps
	Depfile        string
	MSVCDepsPrefix string
	Description    string
	Generator      bool
	In             []string
	InNewline      []string
	Out            []string
	Restat         bool
	Rspfile        string
	RspfileContent string
	Vars           Vars
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
	if r.Depfile != "" {
		writeVar(w2, "depfile", r.Depfile)
	}
	if r.MSVCDepsPrefix != "" {
		writeVar(w2, "msvc_deps_prefix", r.MSVCDepsPrefix)
	}
	if r.Description != "" {
		writeVar(w2, "description", r.Description)
	}
	writeVar(w2, "generator", r.Generator)
	if len(r.In) > 0 {
		writeVar(w2, "in", r.In)
	}
	if len(r.InNewline) > 0 {
		writeVar(w2, "in_newline", r.InNewline)
	}
	if len(r.Out) > 0 {
		writeVar(w2, "out", r.Out)
	}
	writeVar(w2, "restat", r.Restat)
	if r.Rspfile != "" {
		writeVar(w2, "rspfile", r.Rspfile)
	}
	if r.RspfileContent != "" {
		writeVar(w2, "rspfile_content", r.RspfileContent)
	}
	r.Vars.WriteTo(w2)
	return wr.Done()
}

// RequiredVersion returns the ninja version required for the rule
func (r Rule) RequiredVersion() (v Version) {
	// deps since Ninja 1.3
	if r.Deps != 0 {
		v = Ver1_3
	}
	// msvc_deps_prefix since Ninja 1.5
	if r.MSVCDepsPrefix != "" {
		v = Ver1_5
	}
	return
}
