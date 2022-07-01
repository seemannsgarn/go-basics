package main

import (
	"errors"
	"fmt"
)

func main() {
	// messages := [3]string{"1", "2", "3"} // this is the array

	// messages := []string{"a", "b", "c"} // this is the slice

	// messages := make([]string, 5) // initializing a slice with len

	// messages := make([]string, 5) // initializing a slice with len and capacity
	// fmt.Println(len(messages))
	// fmt.Println(cap(messages))

	matrix := make([][]int, 10) // initializating two-dimensional slice

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][x] = x
		}
		fmt.Println(matrix[x])
	}

}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("err: empty array")
	}

	messages[2] = "z" // we can change variable in slice
	fmt.Println(messages)

	return nil
}
