//echo2 output comandline args
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""                  // short define variable
	for _, arg := range os.Args[1:] { // _ is blank indentifier, need because compiler requires variable
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
