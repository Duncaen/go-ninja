package ninja

import (
	"io"

	"github.com/Duncaen/go-ninja/util"
)

// Pool represents a ninja pool block
type Pool struct {
	Name    string
	Depth   int
	Console bool
}

// RequiredVersion returns the ninja version required for the pool
func (p Pool) RequiredVersion() (v Version) {
	// pools are available since Ninja 1.1
	v = Ver1_1
	// console pool is available since Ninja 1.5
	if p.Console {
		v = Ver1_5
	}
	return
}

// WriteTo writes the rule block to w
func (p Pool) WriteTo(w io.Writer) (n int64, err error) {
	wr := util.NewSeqWriter(w)
	io.WriteString(wr, "pool ")
	io.WriteString(wr, p.Name)
	io.WriteString(wr, "\n")
	w2 := util.NewPrefixWriter(wr, []byte("  "))
	writeVar(w2, "depth", p.Depth)
	writeVar(w2, "console", p.Console)
	return wr.Done()
}
