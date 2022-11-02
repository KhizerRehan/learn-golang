package main

import (
	"fmt"
)

func main() {
	// **********************************
	// Bool

	var b bool = true
	var boolTest bool = 1 == 1
	var boolTest2 bool = 1 == 2
	var uninitalizedBool bool

	fmt.Printf("Value: %v, Type: %T", b, b)
	fmt.Println()

	fmt.Printf("Value: %v, Type: %T", boolTest, boolTest)
	fmt.Println()

	fmt.Printf("Value: %v, Type: %T", boolTest2, boolTest2)
	fmt.Println()

	fmt.Printf("Unitialized Boolen Value: %v, Type: %T", uninitalizedBool, uninitalizedBool)
	fmt.Println()

	// **********************************
	// Number

	num1 := 10
	num2 := 10

	fmt.Println("Sum:", num1+num2)
	fmt.Println("Difference:", num1-num2)
	fmt.Println("Multiply:", num1*num2)
	fmt.Println("Divide:", num1/num2)
	fmt.Println("Remainder:", num1%num2)

	// Explicit Conversion
	var num3 int = 10
	var num4 int8 = 3
	// fmt.Println("Not Allowed without conversion:", num3+num4)
	fmt.Println("Explicit Conversion:", num3+int(num4))

}

/*
Summary:
- lower case first letters for package scope
- upper case first letters to export
- block scope
- There is on such thing called private scope

Naming conventions:
- Pascal Case or camelCase
- capitalize acronyms e.g (URL, HTTP)
- keep short and descriptive names
*/
