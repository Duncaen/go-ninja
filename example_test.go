package ninja_test

import (
	"os"

	"github.com/Duncaen/go-ninja"
	"github.com/Duncaen/go-ninja/rule"
	"github.com/Duncaen/go-ninja/build"
)

func ExampleNew() {
	f := ninja.New(
		"build.ninja",
		ninja.Var{"foo", "bar"},
		rule.New("cc",
			rule.DepsGCC,
			rule.Command{"cc -o $out $in"},
			ninja.Var{"foo", "bar"},
		),
		build.New("cc",
			build.ImplicitIn{"bar.h"},
			build.ImplicitOut{"bar.o"},
			build.In{"foo.c"},
			build.OrderOnlyIn{"fizz.h"},
			build.Out{"foo.o"},
			ninja.Var{"foo", "bar"},
		),
	)
	f.WriteTo(os.Stdout)

	// Output:
	// foo = bar
	// rule cc
	//   command = cc -o $out $in
	//   deps = gcc
	//   foo = bar
	// build foo.o | bar.o: cc foo.c | bar.h || fizz.h
	//   foo = bar
}
