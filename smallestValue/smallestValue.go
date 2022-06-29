package main

import "fmt"

func main() {
	fmt.Println(smallest(0, 1, 2, 3, 4, 5))
}

func smallest(numbers ...int) int { // ...int == multiple args
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
