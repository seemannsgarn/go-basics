// map - this is the data structure, a dictionary
package main

import "fmt"

func main() {
	map1()
}
func map1() {
	// users := make(map[string]int) //initializing map
	users := map[string]int{
		"Yoda":           1,
		"Count Dooku":    2,
		"Mace Windu":     3,
		"Qui-Gon Jinn":   4,
		"Depa Billaba":   5,
		"Obi-Wan Kenobi": 6,
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
	// delete element from map
	delete(users, "Qui-Gon Jinn")

}
