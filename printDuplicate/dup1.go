// get stdin "for example from file" and print duplicate lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // make new empty map
	input := bufio.NewScanner(os.Stdin) // get stdin
	for input.Scan() {                  // scanning stdin
		counts[input.Text()]++ // equals: line := input.Text(); counts[line] = counts[line] + 1
	}
	// ignore errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
