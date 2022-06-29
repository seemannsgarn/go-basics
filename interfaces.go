package main

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}
type ShapeWithArea interface {
	Area() float32
}
type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	// square := Square{10}
	// circle := Circle{15}

	// printShapeArea(square)
	// printShapeArea(circle)

	// printInterface(circle)
	// printInterface(square)
	// printInterface("hello")
	// printInterface(22)
	// printInterface(true)

	printInterface("123")
	printInterface(123)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	// switch value := i.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case bool:
	// 	fmt.Println("bool", value)
	// default:
	// 	fmt.Println("unknown type", value)
	// }
	// fmt.Printf("%+v\n", i)

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}
	fmt.Println(len(str))

}
