package main

import "fmt"

func main() {
	defer deferFunc() // call latest in program
	fmt.Println("Hello")

	// msg := []string{
	// 	"line 1",
	// 	"line 2",
	// 	"line 3",
	// 	"line 4",
	// 	"line 5",
	// 	"line 6",
	// }
	// msg[6] = "line 7" // this will be a panic

	// fmt.Println(msg)
	// panic("help me!!!")
}

func deferFunc() {
	fmt.Println("I'am call latest, before exit")
}
func handlerPanic() { // recover panic
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
