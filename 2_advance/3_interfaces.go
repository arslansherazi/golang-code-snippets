package main

import (
	"fmt"
	"math"
)

// interface
type Shape interface {
	calculateArea() float64
}

// structures
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// implementation of interface methods
func (cirlce Circle) calculateArea() float64 {
	return math.Pi * cirlce.radius * cirlce.radius
}

func (rectangle Rectangle) calculateArea() float64 {
	return rectangle.width * rectangle.height
}

//getArea method
func getArea(shape Shape) float64 {
	return shape.calculateArea()
}

// main function
func main() {
	cirlce := Circle{x: 1, y: 2, radius: 5}
	area_of_circle := getArea(cirlce)
	fmt.Println("Area of circle:: ", area_of_circle)

	rectangle := Rectangle{width: 20.2, height: 40.7}
	area_of_rectangle := getArea(rectangle)
	fmt.Println("Area of rectangle:: ", area_of_rectangle)
}
