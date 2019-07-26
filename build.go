package ninja

import (
	"io"

	"github.com/Duncaen/go-ninja/util"
)

// Build reperesents a ninja build block
type Build struct {
	Rule        string
	Out         []string
	OutImplicit []string
	In          []string
	InImplicit  []string
	InOrderOnly []string
	Pool        string
	Vars        Vars
}

// RequiredVersion returns the ninja version required for the build
func (b Build) RequiredVersion() (v Version) {
	// pools since Ninja 1.5
	if b.Pool != "" {
		v = Ver1_5
	}
	// implicit outputs since Ninja 1.7
	if len(b.OutImplicit) > 0 {
		v = Ver1_7
	}
	return
}

// writePaths writes paths to w
func writePaths(w io.Writer, paths []string) {
	len := len(paths)
	for i, p := range paths {
		io.WriteString(w, p)
		if i < len-1 {
			io.WriteString(w, " ")
		}
	}
}

// WriteTo writes the build block to w
func (b Build) WriteTo(w io.Writer) (int64, error) {
	wr := util.NewSeqWriter(w)

	io.WriteString(wr, "build ")
	writePaths(wr, b.Out)

	if len(b.OutImplicit) > 0 {
		io.WriteString(wr, " | ")
		writePaths(wr, b.OutImplicit)
	}

	io.WriteString(wr, ": ")
	io.WriteString(wr, b.Rule)

	if len(b.In) > 0 {
		io.WriteString(wr, " ")
		writePaths(wr, b.In)
	}

	if len(b.InImplicit) > 0 {
		io.WriteString(wr, " | ")
		writePaths(wr, b.InImplicit)
	}

	if len(b.InOrderOnly) > 0 {
		io.WriteString(wr, " || ")
		writePaths(wr, b.InOrderOnly)
	}

	io.WriteString(wr, "\n")
	w2 := util.NewPrefixWriter(wr, []byte("  "))
	if b.Pool != "" {
		writeVar(w2, "pool", string(b.Pool))
	}
	b.Vars.WriteTo(w2)

	return wr.Done()
}

// // Prefix adds a prefix to build path options
// func Prefix(prefix string, v ...interface{}) []interface{} {
// 	var res []interface{}
// 	for _, n := range v {
// 		switch x := n.(type) {
// 		case In:
// 			r := In(util.PrefixPaths(prefix, x...))
// 			res = append(res, r)
// 		case ImplicitIn:
// 			r := ImplicitIn(util.PrefixPaths(prefix, x...))
// 			res = append(res, r)
// 		case OrderOnlyIn:
// 			r := OrderOnlyIn(util.PrefixPaths(prefix, x...))
// 			res = append(res, r)
// 		case ImplicitOut:
// 			r := ImplicitOut(util.PrefixPaths(prefix, x...))
// 			res = append(res, r)
// 		case Out:
// 			r := Out(util.PrefixPaths(prefix, x...))
// 			res = append(res, r)
// 		default:
// 			res = append(res, x)
// 		}
// 	}
// 	return res
// }
