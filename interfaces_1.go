//Go programming provides another data type called interfaces which represents a set of method signatures.
//The struct data type implements these interfaces to have method definitions for the method signature of the interfaces.

package main

import "fmt"

func main() {

	rect := Rectangle{10, 2}
	fmt.Println(executeArea(rect))
}

type Shape interface {
	area() int
}
func executeArea(shape Shape) int {
	return shape.area()
}

type Rectangle struct {
	length, breadth int
}

func (rect Rectangle) area() (result int) {
	result = (rect.length) * (rect.breadth)
	return
}
