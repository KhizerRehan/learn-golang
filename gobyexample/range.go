package main

import (
	"fmt"
)

func main() {
	fmt.Println("Range:")
	// range iterates over elements in a variety of data structures.

	nums := []int{5, 10, 15, 20}
	sum := 0

	for _, num := range nums {
		fmt.Println("Num:", num)
		sum += num
	}
	fmt.Println("Sum", sum)
	fmt.Println()

	// Sometimes we actually want the indexes though.

	for i, num := range nums {
		if num == 20 {
			fmt.Println("Num Value Index is:", i)
		}
	}
	fmt.Println()

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"fname": "khizer", "lname": "rehan"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println()

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}
	fmt.Println()

	// range on strings iterates over Unicode code points.
	for i, c := range "abcde" {
		fmt.Println(i, c)
	}

	fmt.Println()
	for i, c := range "ABCDE" {
		fmt.Println(i, c)
	}
}
