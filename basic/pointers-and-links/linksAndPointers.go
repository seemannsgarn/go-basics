package main

import "fmt"

func main() {
	// var pointer *int //empty pointer
	message := "May the Force be with you! "
	fmt.Println(message)
	changeMessage(&message) // referencing
	fmt.Println(message)
}

func printMessage(message string) { // make new variable
	message += "(from func printMessage)"
	fmt.Println(message)
}

func changeMessage(message *string) { // pointer
	*message += "(from func changeMessage)"

}
