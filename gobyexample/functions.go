package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func multiple(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func minus(a int, b int) int {
	return a - b
}

// When you have multiple consecutive parameters of the same type,
//  you may omit the type name for the like-typed parameters up to
// the final parameter that declares the type

func multipleParameters(a, b, c, d int) int {
	return a + b + c + d
}

func main() {
	fmt.Println("Function")

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = multiple(2, 2)
	fmt.Println("2*2 =", res)

	res = divide(2, 2)
	fmt.Println("2/2 =", res)

	res = minus(2, 2)
	fmt.Println("2-2 =", res)

	res = multipleParameters(2, 4, 6, 8)
	fmt.Println("2+4+6+8 =", res)
}
