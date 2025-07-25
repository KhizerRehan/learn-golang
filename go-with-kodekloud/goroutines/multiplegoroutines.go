package goroutines

import (
	"fmt"
	"time"
)


func CallMultipleGoroutines() {
	for i := 0; i < 10; i++ {
		go startGoroutine()
	}
	time.Sleep(time.Second * 2)
}

func startGoroutine() {
	go inProgress() // this will start a new goroutine there is no parent child relationship
	fmt.Println("Starting goroutine")
}

func inProgress() {
	fmt.Println("In progress")
}