// Copyright 2019 Duncan Overbruck
//
// Permission to use, copy, modify, and/or distribute this software
// for any purpose with or without fee is hereby granted, provided
// that the above copyright notice and this permission notice appear
// in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL
// WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE
// AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL
// DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA
// OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
// TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.
//
// SPDX-License-Identifier: ISC

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
