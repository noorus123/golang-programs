package main

import "fmt"

func main() {

	rect := Rectangle{10, 2}
	fmt.Println("area of rectangle : ", executeArea(rect))

	cir := Circle{3}
	fmt.Println("area of circle : ", executeArea(cir))
}

type Shape interface {
	area() float64
}

func executeArea(shape Shape) float64 {
	return shape.area()
}

type Rectangle struct {
	length, breadth float64
}

func (rect Rectangle) area() (result float64) {
	result = (rect.length) * (rect.breadth)
	return
}

type Circle struct {
	radius float64
}

func (cir Circle) area() float64 {
	return (cir.radius * cir.radius * (3.14))
}
