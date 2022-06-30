package main

import "fmt"

var msg string

func init() { // initialize variables before function main
	msg = "hello youtube"
}
func main() {
	fmt.Println(msg)
}
