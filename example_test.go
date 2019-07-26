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

package ninja_test

import (
	"os"

	"github.com/Duncaen/go-ninja"
)

func ExampleNew() {
	req := &ninja.RequiredVersion{0}
	f := ninja.File{
		ninja.Comment{[]string{"Generated with go-ninja"}},
		req,
		ninja.Var{"foo", "bar"},
		ninja.Pool{"link_pool", 4, false},
		ninja.Pool{"heavy_object_pool", 1, false},
		ninja.Rule{
			Name:    "cc",
			Command: "cc -o $out $in",
			Deps:    ninja.DepsGCC,
			Vars:    []ninja.Var{{"foo", "bar"}},
		},
		ninja.Rule{
			Name:    "link",
			Command: "cc -o $out $in",
		},
		ninja.Build{
			Rule:        "cc",
			In:          []string{"foo.c", ninja.Escape("foo bar.c")},
			InImplicit:  []string{"bar.h"},
			InOrderOnly: []string{"fizz.h"},
			Out:         []string{"foo.o"},
			OutImplicit: []string{"bar.o"},
			Vars:        []ninja.Var{{"foo", "bar"}},
		},
		ninja.Build{
			Rule: "link",
			In:   []string{"foo.o", "bar.o"},
			Out:  []string{"foo"},
			Pool: "link_pool",
		},
	}
	req.Version = f.RequiredVersion()
	f.WriteTo(os.Stdout)

	// Output:
	// # Generated with go-ninja
	// ninja_required_version = 1.7
	// foo = bar
	// pool link_pool
	//   depth = 4
	// pool heavy_object_pool
	//   depth = 1
	// rule cc
	//   command = cc -o $out $in
	//   deps = gcc
	//   foo = bar
	// rule link
	//   command = cc -o $out $in
	// build foo.o | bar.o: cc foo.c foo$ bar.c | bar.h || fizz.h
	//   foo = bar
	// build foo: link foo.o bar.o
	//   pool = link_pool
}
