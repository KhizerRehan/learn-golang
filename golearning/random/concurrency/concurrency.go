package concurrency

import (
	"fmt"
	"time"
)


func printInfinite(s string) {
	for {
		fmt.Println("Infinite", s)
		time.Sleep(1 * time.Millisecond * 100)
	}
}


func CallConcurrencyWithGoRoutine() {
	go printInfinite("Hello") //  launches a new goroutine to run "Hello" printing.
	printInfinite("World")  // runs in the current goroutine (the main one).

	// Since "World" is running in the main goroutine in an infinite loop, the program never exits — and it never reaches the end of CallConcurrencyWithGoRoutine().
}

func CallConcurrencyWithBothGoRoutine() {
	// It will never print anything becasue as soon main function ends, the program exits.
	// Both method are running in separate goroutines, so they are not waiting for each other.
	// So, the program will exit before the goroutines have a chance to print anything.
	go printInfinite("Hello")
	go printInfinite("World")
}


func CallConcurrency() {
	printInfinite("Hello")
	printInfinite("World") // This will not be executed because the main function will exit after the first printInfinite call
}