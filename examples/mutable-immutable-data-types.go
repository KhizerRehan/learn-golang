package main

import (
	"fmt"
)

func changeFirst(slice []int) {
	slice[0] = 5000
}

func main() {
	fmt.Println("Learn Mutable and Immutable DataTypes")

	// It Mutates original reference Slice Datatype

	fmt.Println("Mutable datatype: Slice")
	var x []int = []int{3, 4, 5}
	y := x
	y[0] = 100
	fmt.Println(x, y)
	fmt.Println()

	// It Mutates original reference Slice Map
	fmt.Println()
	fmt.Println("Mutable datatype: Map")
	var mapVar map[string]int = map[string]int{"hello": 3}
	newVar := mapVar
	newVar["y"] = 100
	newVar["7"] = 200
	fmt.Println(mapVar, newVar)

	// It creates a new COPY of array
	fmt.Println()
	fmt.Println("Array datatype: Array")
	var array [2]int = [2]int{1, 3}
	newArray := array
	newArray[0] = 500
	newArray[1] = 1000
	fmt.Println(array, newArray)

	// Mutates by Reference inside method as it directly modifying actual reference of memory
	fmt.Println()
	var dummyArray []int = []int{10, 20, 30}
	fmt.Println("Before", dummyArray)
	changeFirst(dummyArray)
	fmt.Println("After", dummyArray)
}
