// structure - custom type, like a class
package main

import "fmt"

type Car struct {
	vendor string
	model  string
	color  string
	year   int
}

// value receiver
func (c Car) printCarInfoVR() {
	fmt.Println(c.vendor, c.model, c.color, c.year)
}

// pointer receiver
func (c *Car) printCarInfoPR() {
	fmt.Println(c.vendor, c.model, c.color, c.year)
}

func NewCar(vendor, model, color string, year int) Car {
	return Car{
		vendor: vendor,
		model:  model,
		year:   year,
		color:  color,
	}
}

func main() {
	car1 := NewCar("Ford", "Mustang", "Black", 1969)
	car2 := NewCar("Toyota", "Celica", "White", 1995)
	car3 := NewCar("Subaru", "Impreza", "Blue", 2003)

	car1.printCarInfoVR()
	car2.printCarInfoVR()
	car3.printCarInfoVR()

	car2.printCarInfoPR()

}
