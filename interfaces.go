package main

import (
	"fmt"
	"math"
)
type ShapeOperation interface{
	getArea() float32 
	getPerimeter() float32
}

type Rectangle struct{
	length float32 
	breadth float32
}

func (rect Rectangle) getArea()float32{
	return rect.length * rect.breadth
}

func (rect Rectangle) getPerimeter() float32{
	return 2 * (rect.length + rect.breadth)
}

type Circle struct{
	radius float32
}

func (circ Circle) getArea()float32{
	return math.Pi * circ.radius * circ.radius
}

func (circ Circle) getPerimeter()float32{
	return 2 * math.Pi * circ.radius
}

func makeShapes(){
	rectangle := Rectangle{length: 10, breadth: 5}
	fmt.Println("Area of rect:", rectangle.getArea())
	fmt.Println("Perimeter of rect: ", rectangle.getPerimeter())

	circle := Circle{radius: 7}
	fmt.Println("Area of circle: ", circle.getArea())
	fmt.Println("Perimeter of circle: ", circle.getPerimeter())
}