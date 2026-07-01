package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	p1 := Person{
		name: "Khizer Rehan",
		age:  32,
	}

	fmt.Print("Person Details", p1)
}
