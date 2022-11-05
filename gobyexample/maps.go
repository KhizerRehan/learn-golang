package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learn Maps")

	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	hashMap := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	hashMap["key1"] = 10
	hashMap["key2"] = 20
	fmt.Println("HashMap:", hashMap)

	// Get a value for a key with name[key].
	keyValue := hashMap["key1"]
	fmt.Println("Single Key =>", keyValue)

	// The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("Length:", len(hashMap))

	// The builtin delete removes key/value pairs from a map.
	fmt.Println("Delete Key2 from HashMap")
	delete(hashMap, "key2")
	fmt.Println("HashMap:", hashMap)

	_, prs := hashMap["key2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"k1": 10, "k2": 20}
	fmt.Println("map:", n)
}
