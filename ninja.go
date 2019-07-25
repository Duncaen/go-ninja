// Package ninja implements a ninja build file generator
package ninja

import (
	"io"

	"github.com/Duncaen/go-ninja/util"
)

// Node represents a ninja variable, statement or block
type Node interface {
	WriteTo(wr io.Writer) (n int64, err error)
}

// File represents a ninja build file
type File struct {
	Name string
	Nodes []Node
}

// New creates a new ninja file
func New(name string, nodes ...Node) File {
	return File{name, nodes}
}

// WriteTo writes the complete ninja file to w
func (f File) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)
	for _, n := range f.Nodes {
		n.WriteTo(wr)
	}
	return wr.Done()
}
