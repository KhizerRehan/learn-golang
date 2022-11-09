package main

import (
	"fmt"
)

/*
This function intSeq returns another function,
which we define anonymously in the body of intSeq.
The returned function closes over the variable i to form a closure.
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fullName(fname string) func(l string) string {
	return func(lname string) string {
		v := fmt.Sprintf("%s %s", fname, lname)
		return v
	}

}

func main() {
	fmt.Println("Closure Example")

	nextInt := intSeq()
	fmt.Println("Next Int", nextInt())
	fmt.Println("Next Int", nextInt())
	fmt.Println("Next Int", nextInt())
	fmt.Println("Next Int", nextInt())
	fmt.Println("Next Int", nextInt())
	fmt.Println("Next Int", nextInt())
	fmt.Println()

	name := fullName("Khizer")
	completeName := name("Rehan")
	fmt.Println("Full Name:", completeName)
}
