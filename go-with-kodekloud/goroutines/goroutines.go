package goroutines

import (
	"fmt"
	"time"
)


func CallGoroutines() {
	fmt.Println("CallGoroutines")

	start := time.Now()
	for i := 0; i < 10000; i++ {
		go CalculateSequential(i)
	}	
	elapsed := time.Since(start)
	time.Sleep(time.Second * 2)
	fmt.Println("Time taken:", elapsed)
}

func CalculateSequential(a int) {
	fmt.Println("Square of", a, "is", a*a)
}