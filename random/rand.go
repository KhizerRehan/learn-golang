package main

import "fmt"


type Counter struct {
	value int
}


func (c *Counter) increment() {
	c.value++
}

func incrementCounter(c *Counter) {
	c.value++
}


func main() {

	counter := Counter{value: 0}

	fmt.Println("Before: Received Func", counter.value)
	counter.increment()
	fmt.Println("After: Received Func", counter.value)



	

	counter1 := Counter{value: 0}
	fmt.Println("Before: Passed Func", counter1.value)
	incrementCounter(&counter1)
	fmt.Println("After: Passed Func", counter1.value)	

	
}	