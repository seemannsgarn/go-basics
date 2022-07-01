package main

import "fmt"

func main() {
	msg := []string{
		"line 1",
		"line 2",
		"line 3",
		"line 4",
		"line 5",
		"line 6",
	}
	for i := range msg { // equals = for i :0; i < len(msg); i++, we can get index(_) and value
		fmt.Println(msg[i])
	}
	for _, value := range msg {
		fmt.Println("value =", value)
	}
}
