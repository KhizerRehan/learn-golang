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

func (p Person) details() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

func main() {
	fmt.Println("Hello World")

	p1 := NewPerson("Khizer Rehan", 32)
	p2 := NewPerson("John Doe", 25)

	fmt.Print("Person Details", p1.details())
	fmt.Println()
	fmt.Print("Person Details", p2.details())
	fmt.Println()

}
