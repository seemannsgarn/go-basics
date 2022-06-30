package main

import "fmt"

func main() {
	fmt.Println(enterTheSite(50))
}

// func print(arg string) {
// 	fmt.Println(arg)
// }

// func sayHello1(name string) string {
// 	return "Hello " + name + "!"
// }

// func sayHello2(name string, age int) string {
// 	return fmt.Sprintf("Hello, %s! Your %d old! ", name, age)
// }

func enterTheSite(age int) (string, bool) {
	if age >= 18 && age <= 45 { // && == and, || == or
		return "18-45", true
	} else if age >= 45 && age < 65 {
		return "45-65", true
	} else if age >= 65 {
		return "65+", false
	}
	return "18-", false
}
