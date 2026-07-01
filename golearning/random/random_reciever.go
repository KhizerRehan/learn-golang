package random

import "fmt"

type Calculator struct {
	val1 int
	val2 int
}

// reciever is a function that is associated with a type

func (c Calculator) Add() int {
	return c.val1 + c.val2
}

func (c Calculator) Subtract() int {
	return c.val1 - c.val2
}

func (c Calculator) Multiply() int {
	return c.val1 * c.val2
}

func NewCalculator(val1 int, val2 int) *Calculator {
	return &Calculator{
		val1: val1,
		val2: val2,
	}
}

func CallCalculator() {

	calculator := Calculator{
		val1: 10,
		val2: 10,
	}

	fmt.Println("Addition:", calculator.Add())
	fmt.Println("Subtraction:", calculator.Subtract())
	fmt.Println("Multiplication:", calculator.Multiply())


	// Using NewCalculator
	fmt.Println("Using NewCalculator")
	calculator2 := NewCalculator(20, 20)
	fmt.Println("Addition:", calculator2.Add())
	fmt.Println("Subtraction:", calculator2.Subtract())
	fmt.Println("Multiplication:", calculator2.Multiply())
}
