package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Define a Wrapper Mehthod

func calculateArea(s Shape) float64 {
	/*
		The calculateArea function accepts a Shape interface, allowing any implementing type to be passed.
		When called with rectangle or circle, the respective variables are implicitly converted to a Shape interface, 
		invoking the Area method on the underlying struct.

		Imp!!
		As Rectangle and Circle both implement the Shape interface, they can be passed to the calculateArea function.
	*/

	return s.Area()
}

func calculatePerimeter(s Shape) float64 {
	return s.Perimeter()
}


func CallInterfaceV1() {
	// Using the struct (Rectangle) directly, not assigning it to a variable of type Shape.
	// Go does not require a struct to implement an interface unless you explicitly use the interface type.

	// Directly Calling Struct
	rectangle := Rectangle{Width: 10, Height: 10}
	fmt.Printf("Area of Rectangle with width %v and height %v is %v\n", rectangle.Width, rectangle.Height, rectangle.Area())
	circle := Circle{Radius: 10}
	fmt.Printf("Area of Circle with radius %v is %v\n", circle.Radius, circle.Area())


	// Calling Using Wrapper Method
	fmt.Println("Rectangle Area:", calculateArea(rectangle))
	fmt.Println("Circle Area:", calculateArea(circle))
	fmt.Printf("Rectangle Perimeter: %.2f\n", calculatePerimeter(rectangle))
	fmt.Printf("Circle Perimeter: %.2f\n", calculatePerimeter(circle))

}
