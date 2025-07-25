// package random

// import (
// 	"fmt"
// )

// // ChannelCollector is a struct that collects random numbers from a channel

// type ChannelCollector struct {
// 	collect chan int
// }

// func NewChannelCollector() *ChannelCollector {
// 	return &ChannelCollector{
// 		collect: make(chan int),
// 	}
// }

// func (c *ChannelCollector) Collect(n int) {
// 	c.collect <- n
// }

// func CallChannelCollector() {
// 	c := NewChannelCollector()

// 	go func() {
// 		c.Collect(10)
// 		c.Collect(20)
// 		c.Collect(30)
// 	}()

// 	for i := 0; i < 3; i++ {
// 		fmt.Println(<-c.collect)
// 	}
// }

package random

import (
	"fmt"
	"sync"
	"time"
)

type ChannelCollector struct {
	collect chan int
	wg      sync.WaitGroup
}

func NewChannelCollector() *ChannelCollector {
	return &ChannelCollector{
		collect: make(chan int),
	}
}

func (c *ChannelCollector) Collect(n int) {
	fmt.Println("Adding work: wg.Add(1)")
	c.wg.Add(1)

	go func(num int) {
		defer func() {
			fmt.Println("Work done: wg.Done()", num)
			c.wg.Done()
		}()

		// Simulate some work
		time.Sleep(time.Millisecond * 500)
		c.collect <- num
	}(n)
}

func (c *ChannelCollector) WaitAndClose() {
	go func() {
		fmt.Println("Waiting for all goroutines to finish...")
		c.wg.Wait()
		fmt.Println("All work done. Closing channel.")
		close(c.collect)
	}()
}

func CallChannelCollector() {
	c := NewChannelCollector()

	// Launch multiple goroutines
	c.Collect(10)
	c.Collect(20)
	c.Collect(30)

	// Wait for all goroutines and close channel
	c.WaitAndClose()

	// Range over channel until it's closed
	for val := range c.collect {
		fmt.Println("Received:", val)
	}
}