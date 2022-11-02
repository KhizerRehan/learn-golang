package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello" + "World")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println("\n AND Operators\n")
	fmt.Println("true && true=", true && true)
	fmt.Println("true && false=", true && false)
	fmt.Println("false && false=", false && false)
	fmt.Println("false && false=", false && false)

	fmt.Println("\nOR Operators\n")
	fmt.Println("true  || true=", true || true)
	fmt.Println("true  || false=", true || false)
	fmt.Println("false || false=", false || false)
	fmt.Println("false || false=", false || false)

	fmt.Println("\nNegation\n")
	fmt.Println("!true=", !true)
	fmt.Println("!true=", !false)
}
