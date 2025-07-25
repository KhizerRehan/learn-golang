package main

import (
	"fmt"
	"learn-golang/go-with-kodekloud/exercises/refresh"
	"learn-golang/go-with-kodekloud/goroutines"
)


func main() {
	fmt.Println("Hello World")

	// callRefresh()
	// callGoroutines()
	callMultipleGoroutines()
}

func callRefresh() {
	fmt.Println(refresh.Calculate(20, 10))
	fmt.Println(refresh.Calculate(700, 70))
}

func callGoroutines() {
	goroutines.CallGoroutines()
}

func callMultipleGoroutines() {
	goroutines.CallMultipleGoroutines()
}