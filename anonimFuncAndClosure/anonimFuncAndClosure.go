package main

import "fmt"

func main() {
	// func() {
	// 	fmt.Println("I'am a closure.")
	// }()
	inc := increment() // make function inc and save state
	fmt.Println(inc())
}

func increment() func() int {
	count := 0 // count in func increment
	return func() int {
		count++
		return count
	}
}
