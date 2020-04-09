/*
Package gobank implements a simple bank account system.
The logic is based on an event sourced model.

Usage:

    ./hello
*/
package main // executable commands must always use package main.

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
