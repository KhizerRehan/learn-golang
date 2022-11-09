package main

import (
	"fmt"
)

// function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Total Sume is %v\n", total)
}

func main() {
	fmt.Println("Variadic Function Examples")
	sum(5, 10)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
