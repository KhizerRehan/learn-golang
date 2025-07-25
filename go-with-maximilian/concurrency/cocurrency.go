package concurrency

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RunTasksInParallel demonstrates concurrent execution using goroutines and channels
// It takes a channel parameter to signal when all tasks are complete
func RunTasksInParallel(doneChan chan bool) {
	// Launch first goroutine - runs concurrently with other tasks
	// The 'go' keyword makes this function execute asynchronously
	go func() {
		time.Sleep(time.Second * 2) // Simulate 2 seconds of work
		fmt.Println("Task 1 is done")
	}()

	// Launch second goroutine - also runs concurrently
	go func() {
		time.Sleep(time.Second * 3) // Simulate 3 seconds of work
		fmt.Println("Task 2 is done")
	}()

	// Launch third goroutine - this one signals completion
	go func() {
		time.Sleep(time.Second * 4) // Simulate 4 seconds of work (longest task)
		fmt.Println("Task 3 is done")
		// Send signal through channel to indicate all tasks are complete
		// Since Task 3 takes the longest (4 seconds), it will finish last
		doneChan <- true
	}()
	
	// Note: All three tasks run in parallel, not sequentially
	// Total execution time will be ~4 seconds (not 2+3+4=9 seconds)
	// The calling function should use <-doneChan to wait for completion
}



func RunTasksInParallelWithMultipleChannels(doneChan chan bool, doneChan2 chan bool) {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Task 1 is done")
		doneChan <- true
	}()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Task 2 is done")
		doneChan2 <- true
	}()
}


func RunTasksInParallelWithWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Task 1 is done")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Task 2 is done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("RunTasksInParallelWithWaitGroup: Tasks are done")
	
}

func RunTaskInParallelWithErrorChannel(doneChan chan bool, errChan chan error) {

	// Generate a Error instead of using done use errChan
	go func() {
		time.Sleep(time.Second * 2)
		// Simulate an error condition (50% chance of error)
		randomValue := rand.Intn(2)
		fmt.Println("Random value:", randomValue)
		if randomValue == 0 {
			errChan <- errors.New("task 1 failed")
			return
		}
		fmt.Println("Task 1 is done")
		doneChan <- true
	}()
		
}

func RunTaskWithInterfaceChannel(doneChan []chan interface{}) {
	go func() {
		time.Sleep(time.Second * 2)
		doneChan[0] <- "khizer"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		doneChan[1] <- "ali"
	}()

	go func() {
		time.Sleep(time.Second * 4)
		doneChan[2] <- "ahmed"
	}()
}


type Result struct {
    Data  interface{}
    Error error
}


func RunTaskWithErrorChannel(doneChan []chan Result) {
    go func() {
        time.Sleep(time.Second * 2)
        doneChan[0] <- Result{Data: "khizer", Error: nil}
    }()

    go func() {
        time.Sleep(time.Second * 3)
        doneChan[1] <- Result{Data: nil, Error: fmt.Errorf("something went wrong")}
    }()
}