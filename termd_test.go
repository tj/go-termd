package termd_test

import (
	"fmt"
	"io/ioutil"

	"github.com/tj/go-termd"
)

func Example() {
	var c termd.Compiler

	c.SyntaxHighlighter = termd.SyntaxTheme{
		"keyword": termd.Style{},
		"comment": termd.Style{
			Color: "#323232",
		},
		"literal": termd.Style{
			Color: "#555555",
		},
		"name": termd.Style{
			Color: "#777777",
		},
		"name.function": termd.Style{
			Color: "#444444",
		},
		"literal.string": termd.Style{
			Color: "#333333",
		},
	}

	b, _ := ioutil.ReadFile("testdata/example.md")
	fmt.Printf("%s\n", c.Compile(string(b)))
	// Output:
}
