// Package ninja implements a ninja build file generator
package ninja

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/Duncaen/go-ninja/util"
)

// Version represents the ninja version requirement
type Version int

const (
	// Ver1_1 to require ninja 1.1
	Ver1_1 Version = iota + 1
	// Ver1_3 to require ninja 1.3
	Ver1_3
	// Ver1_5 to require ninja 1.5
	Ver1_5
	// Ver1_7 to require ninja 1.7
	Ver1_7
)

func (v Version) String() string {
	switch v {
	case Ver1_3:
		return "1.3"
	case Ver1_5:
		return "1.5"
	case Ver1_7:
		return "1.7"
	}
	return ""
}


// Node represents a ninja variable, statement or block
type Node interface {
	WriteTo(wr io.Writer) (n int64, err error)
	RequiredVersion() Version
}

// File represents a ninja build file
type File []Node

// RequiredVersion returns the ninja version for the file
func (f File) RequiredVersion() Version {
	var res Version
	for _, n := range f {
		v := n.RequiredVersion()
		if v > res {
			res = v
		}
	}
	return res
}

// WriteTo writes the complete ninja file to w
func (f File) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)
	for _, n := range f {
		n.WriteTo(wr)
	}
	return wr.Done()
}

// Escape escapes special characters like ":" to become "$:".
// It escapes the following characters: "$", ":", " " (space), "\n".
//
// https://ninja-build.org/manual.html#_lexical_syntax
func Escape(str string) string {
	var b bytes.Buffer
	for _, c := range str {
		switch c {
		case '$', ':', ' ', '\n':
			b.WriteRune('$')
			b.WriteRune(c)
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

// writesVar writes a variable definition to w
func writeVar(w io.Writer, key string, v interface{}) {
	var val string
	switch x := v.(type) {
	case bool:
		if !x {
			return
		}
		val = "true"
	case int:
		val = strconv.Itoa(x)
	case string:
		val = x
	case []string:
		val = strings.Join(x, " ")
	}
	io.WriteString(w, key)
	io.WriteString(w, " = ")
	io.WriteString(w, val)
	io.WriteString(w, "\n")
}

// Comment represents a ninja file comment
type Comment struct {
	Lines []string
}

// WriteTo writes the comment to w
func (c Comment) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(util.NewPrefixWriter(w, []byte("# ")))
	for _, line := range c.Lines {
		io.WriteString(wr, line)
		io.WriteString(wr, "\n")
	}
	return wr.Done()
}

// RequiredVersion returns Version(0), since comments were always supported
func (c Comment) RequiredVersion() Version {
	return Version(0)
}

// RequiredVersion represents the ninja_required_version variable
type RequiredVersion struct {
	Version Version
}

// WriteTo writes the required version to w
func (r RequiredVersion) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)
	if r.Version > Version(0) {
		writeVar(wr, "ninja_required_version", r.Version.String())
	}
	return wr.Done()
}

// RequiredVersion returns the required version
func (r RequiredVersion) RequiredVersion() Version {
	return r.Version
}
