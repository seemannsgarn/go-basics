package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, err := checkNum(-55)
	if err != nil {
		log.Println(err) // another package for logs
		return
	}

	fmt.Println(message)
}

func checkNum(age int) (string, error) {
	if age == 0 {
		return "Zero", errors.New("0 can't be")
	} else if age >= 1 {
		return "It's a fine number", nil // nil == ignore err
	} else {
		return "negative num", errors.New("wtf")
	}
}
