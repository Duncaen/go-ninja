package rule_test

import (
	"os"

	"github.com/Duncaen/go-ninja"
	"github.com/Duncaen/go-ninja/rule"
)

func ExampleNew() {
	r := rule.New("cc",
		rule.DepsGCC,
		rule.Command{"cc -o $out $in"},
		ninja.Var{"foo", "bar"},
	)
	r.WriteTo(os.Stdout)

	// Output:
	// rule cc
	//   command = cc -o $out $in
	//   deps = gcc
	//   foo = bar
}
