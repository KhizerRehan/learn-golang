package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		name,
		age,
	}
}

func main() {

	p1 := NewPerson("Khizer Rehan", 32)
	p2 := NewPerson("John Doe", 25)

	fmt.Print("Person Details", p1)
	fmt.Println()
	fmt.Print("Person Details", p2)
	fmt.Println()

}
