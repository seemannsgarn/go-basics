//echo1 output comandline args
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // define variables s and sep with type string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
