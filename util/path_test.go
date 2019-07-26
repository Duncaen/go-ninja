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
)

var testPrefixPaths = []struct {
	prefix string
	paths  []string
	out    []string
}{
	{
		prefix: "foo",
		paths:  []string{"bar"},
		out:    []string{"foo/bar"},
	},
	{
		prefix: "fizz/buzz",
		paths:  []string{"foo/bar"},
		out:    []string{"fizz/buzz/foo/bar"},
	},
	{
		prefix: "fizz/buzz",
		paths:  []string{"foo", "bar"},
		out:    []string{"fizz/buzz/foo", "fizz/buzz/bar"},
	},
}

func testEQPaths(t *testing.T, got, exp []string) bool {
	t.Helper()
	if len(got) != len(exp) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if got[i] != exp[i] {
			return false
		}
	}
	return true
}

func TestPrefixPaths(t *testing.T) {
	for _, tt := range testPrefixPaths {
		ret := PrefixPaths(tt.prefix, tt.paths...)
		if !testEQPaths(t, ret, tt.out) {
			t.Errorf("expected %v, got %v", tt.out, ret)
		}
	}
}
