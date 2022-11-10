package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name.
func newPerson(name string) *Person {

	// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println("Struct Examples")

	// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Bob", age: 30})
	// Omitted fields will be zero-valued.
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println()

	fmt.Println("New Person")
	fmt.Println(newPerson("Jon"))

	p1 := newPerson("Khizer Rehan")
	fmt.Println("&p", &p1)
	fmt.Println()
	// ----------------------------
	// BOTH ARE SAME

	fmt.Println("p", p1)
	fmt.Println("*&p", (*&p1)) // *(defrenced) -> &address
	fmt.Println()
	// ----------------------------
	// BOTH ARE SAME

	fmt.Println("*p", *p1)
	fmt.Println("**&p", (**&p1)) //  -> *(defrenced) -> *(defrenced) -> &address
	fmt.Println()
	// ----------------------------
	// You can also use dots with struct pointers - the pointers are automatically dereferenced

	fmt.Println("p1.name", p1.name)
	fmt.Println("p1.age", p1.age)
	fmt.Println()

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	// Update Reference Memory Value
	fmt.Println("Structs are mutable by reference")
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(sp.age)
}
