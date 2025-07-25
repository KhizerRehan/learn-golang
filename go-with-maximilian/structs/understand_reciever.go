package structs

import (
	"errors"
	"fmt"
)

type Counter struct {
	value int
}


func (c *Counter) Increment() {
	(*c).value++
}

func (c *Counter) Decrement() {
	(*c).value--
}

func (c Counter) GetValue() int {
	return c.value
}

func NewCounter(value int) (*Counter, error) {

	if value < 0 {
		return nil, errors.New("value must be greater than 0")
	}

	return &Counter{value: value}, nil
}


func CounterWrapper() {
	counter, err := NewCounter(10)

	if err != nil {
		fmt.Println("Error creating counter:", err)
		return
	}

	counter.Increment()
	fmt.Println("Counter value:", counter.GetValue())
	counter.Decrement()
	fmt.Println("Counter value:", counter.GetValue())

	counter2, err2 := NewCounter(-5)

	if err2 != nil {
		fmt.Println("Error creating counter:", err2)
		return
	}

	counter2.Increment()
	fmt.Println("Counter value:", counter2.GetValue())
	counter2.Decrement()
	fmt.Println("Counter value:", counter2.GetValue())

	
}