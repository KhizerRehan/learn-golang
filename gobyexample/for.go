package main

import (
	"fmt"
)

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 5; j <= 9; j++ {
		fmt.Println("Index: ", j)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	var oddSum int = 0
	var evenSum int = 0
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println("Even Value: ", n)
			evenSum += n
		} else {
			fmt.Println("Odd Value: ", n)
			oddSum += n
		}
	}
	fmt.Println("\nOdd Sum:", oddSum)
	fmt.Println("Even Sum:", evenSum)

}
