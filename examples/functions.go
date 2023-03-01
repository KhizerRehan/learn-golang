package main

import (
	"fmt"
)

// Functions

func test() {
	fmt.Println("Test function")
}

func acceptingMultiplyFunc(myFunc func(int) int) {
	fmt.Println(myFunc(10))
}

func acceptingAddFunc(myFunc func(int, int) int) {
	fmt.Println(myFunc(50, 50))
}

func returnClosure(value string) func() {
	return func() {
		fmt.Println(value)
	}
}

func fullNameClosure(fname string) func(string) {
	return func(lname string) {
		fmt.Println(fname, lname)
	}
}

func main() {
	test()
	x := test
	x()

	// function inside a function
	storingfunc := func(x int) int {
		return x * x
	}(5) // # immediatly invoking a function think like IIFE

	fmt.Println(storingfunc)

	// passing a function to another function

	multiplyFunction := func(x int) int {
		return x * x
	}

	addFunction := func(x int, y int) int {
		return x + y
	}

	acceptingMultiplyFunc(multiplyFunction)
	acceptingAddFunc(addFunction)

	// IIFE fucntion

	func() {
		fmt.Println("Hello, IIFE. I am just comapring with JS")
	}()

	// Closure Example

	returnClosure("Hello")()
	returnClosure("Closure")()
	fullNameClosure("Khizer")("Rehan")
}
