package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays Example")

	//  Here we create an array a that will hold exactly 5 ints.
	//  The type of elements and length are both part of the array’s type
	//  By default an array is zero-valued, which for ints means 0s.

	var arrEmp [5]int
	fmt.Println("Emp: ", arrEmp)

	// We can set a value at an index using the array[index] = value syntax, and get a value with array[index].

	arrEmp[4] = 100
	fmt.Println("Value Setted at which index => Set", arrEmp)
	fmt.Println("Value Retrived at particular index => Get", arrEmp)
	fmt.Println("Length of Array", len(arrEmp))

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

	var twoDimensional [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Row:Column => (%v %v)\n", i, j)
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("twoDimensional: ", twoDimensional)

}
