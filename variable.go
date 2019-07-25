package ninja

import (
	"io"

	"github.com/Duncaen/go-ninja/util"
)

// Var represents a ninja variable
type Var struct {
	Key, Val string
}

// WriteTo writes the variable to w
func (v Var) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)
	io.WriteString(wr, v.Key)
	io.WriteString(wr, " = ")
	io.WriteString(wr, v.Val)
	io.WriteString(wr, "\n")
	return wr.Done()
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
