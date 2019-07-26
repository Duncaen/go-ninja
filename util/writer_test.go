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

package util

import (
	"testing"
	"bytes"
)

var prefixTests = []struct {
	pre string
	in []string
	out string
}{
	{">>", []string{"a = b\n"}, ">>a = b\n"},
	{">>", []string{"a = b\n", "c = d\n"}, ">>a = b\n>>c = d\n"},
	{">>", []string{"a =", " b\n"}, ">>a = b\n"},
	{">>", []string{"a", " ", "=", " ", "b", "\n"}, ">>a = b\n"},
	{">>", []string{"\n"}, ">>\n"},
	{">>", []string{"\n\n"}, ">>\n>>\n"},
	{">>", []string{"\n", "\n"}, ">>\n>>\n"},
	{">>", []string{"\n", "a"}, ">>\n>>a"},
}

func TestPrefixWriter(t *testing.T) {
	for _, tt := range prefixTests {
		var b bytes.Buffer
		wr := NewPrefixWriter(&b, []byte(tt.pre))
		for _, s := range tt.in {
			wr.Write([]byte(s))
		}
		res := b.String()
		if res != tt.out {
			t.Errorf("expected %q got %q", tt.out, res)
		}
	}
}
