package main

import (
	"fmt"
)

func main() {
	var a = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "appled"
	fmt.Println(f)
	fmt.Printf("Value: %v, Type: %T", f, f)
}
