package ninja

import (
	"io"

	"github.com/Duncaen/go-ninja/util"
)

// Var represents a ninja variable
type Var struct {
	Key string
	Val interface{}
}

// WriteTo writes the variable to w
func (v Var) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)
	writeVar(wr, v.Key, v.Val)
	return wr.Done()
}

// RequiredVersion returns Version(0), since variables were always supported
func (v Var) RequiredVersion() Version {
	return Version(0)
}

// Vars represents multiple variables
type Vars []Var

// WriteTo writes all variables to w
func (vs Vars) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)
	for _, v := range vs {
		v.WriteTo(wr)
	}
	return wr.Done()
}

// RequiredVersion returns Version(0), since variables were always supported
func (vs Vars) RequiredVersion() Version {
	return Version(0)
}
