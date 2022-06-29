package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(checkDrink("vodka"))
}

func checkDrink(drink string) (string, error) {
	switch drink {
	case "whisky":
		return "You need a cigare", nil
	case "rome":
		return "You need a ship", nil
	case "wine":
		return "You need a love", nil
	case "beer":
		return "You need a toilet", nil
	case "vodka":
		return "You need a balalayka and a bear, suka blyat", nil
	default:
		return "You need a better drink", errors.New("invalid drink")
	}
}
