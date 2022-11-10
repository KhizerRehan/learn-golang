package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Strings and Runes")

	// The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
	const str = "Khizer"
	fmt.Println("Str Length", len(str))

	fmt.Println()

	// This loop generates the hex (%x) values of all the bytes that constitute the code points in str and (%v -> Value ASCII)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x, %v", str[i], str[i])
		fmt.Println()
	}

	// In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune
	fmt.Println("Rune count:", utf8.RuneCountInString(str))

	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
