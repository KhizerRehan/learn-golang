package main

import (
	"fmt"
)

// PACKAGE SCOPE
var (
	firstName string = "Khizer"
	lastName  string = "Rehan"
	fullName  string = "Khizer Rehan"
)

var shadowingVariable int = 50

func main() {
	//  FUNCTIONAL SCOPE || BLOCK SCOPE

	var i int
	i = 10
	fmt.Println("Value is:", i)
	fmt.Println()

	var j int = 27
	fmt.Printf("Value: %v, Type: %T", j, j)
	fmt.Println()

	k := 10.12
	fmt.Printf("Value: %v, Type: %T", k, k)
	fmt.Println()

	l := "khizer rehan"
	fmt.Printf("Value: %v, Type: %T", l, l)
	fmt.Println()

	// **********************************
	// Shahdowing
	fmt.Printf("No Shahdowing value: %v, Type: %T", shadowingVariable, shadowingVariable)
	fmt.Println()

	var shadowingVariable int = 100
	fmt.Printf("Shahdowing value: %v, Type: %T", shadowingVariable, shadowingVariable)
	fmt.Println()

	// **********************************
	// Type Casting

	var castingInt int = 42
	var castFloat float32

	castFloat = float32(castingInt)

	fmt.Println("Casting Variabled")
	fmt.Printf("Value: %v, Type: %T", castingInt, castingInt)
	fmt.Println()
	fmt.Printf("Value: %v, Type: %T", castFloat, castFloat)
	fmt.Println()

	var castingFloat float64 = 42.5
	var castToInt int
	castToInt = int(castFloat)

	fmt.Printf("Value: %v, Type: %T", castingFloat, castingFloat)
	fmt.Println()
	fmt.Printf("IMPORTTANT YOU LOST FRACTIONAL PART: Value: %v, Type: %T", castToInt, castToInt)
	fmt.Println()

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
