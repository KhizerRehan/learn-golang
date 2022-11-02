package main

import (
	"fmt"
	"math"
)

const str string = "Some Constant Value"

func main() {

	fmt.Printf("Const Value: %s, Type: %T", str, str)
	const n = 50000
	fmt.Println()
	fmt.Println(math.Sin(n))

	const d = n / 10000
	fmt.Println(d)
	fmt.Println()
}
