package build_test

import (
	"os"

	"github.com/Duncaen/go-ninja"
	"github.com/Duncaen/go-ninja/build"
)

func ExampleBuild() {
	b := build.New("cc",
		build.ImplicitIn{"bar.h"},
		build.ImplicitOut{"bar.o"},
		build.In{"foo.c"},
		build.OrderOnlyIn{"fizz.h"},
		build.Out{"foo.o"},
		ninja.Var{"foo", "bar"},
	)
	b.WriteTo(os.Stdout)

	// Output:
	// build foo.o | bar.o: cc foo.c | bar.h || fizz.h
	//   foo = bar
}
