package structs

import (
	"fmt"
	"strconv"
	"time"
	"github.com/Pallinder/go-randomdata"
)

type Person struct {
	Name string
	Age int
	Email string
	CreatedAt time.Time
}

func BasicStruct() {

	person := Person{
		Name: randomdata.SillyName(),
		Age: randomdata.Number(1, 100),
		Email: randomdata.Email(),
		CreatedAt: time.Now(),
	}

	fmt.Println(person)
	fmt.Println(person.Name +" "+ strconv.Itoa(person.Age) + " " + person.Email)
	fmt.Printf("%s %d %s\n", person.Name, person.Age, person.Email)
	fmt.Println("--------------------------------")

	var personA Person
	personA.Name = randomdata.SillyName()
	personA.Age = randomdata.Number(1, 100)
	personA.Email = randomdata.Email()
	personA.CreatedAt = time.Now()

	fmt.Println(personA)
	fmt.Println(personA.Name +" "+ strconv.Itoa(personA.Age) + " " + personA.Email)
	fmt.Printf("%s %d %s\n", personA.Name, personA.Age, personA.Email)

	fmt.Println("--------------------------------")
	
	// Empty Struct
	var emptyPerson Person
	fmt.Println(emptyPerson)

	// Anonymous Struct
	personB := Person{}
	fmt.Println(personB)

	fmt.Println("--------------------------------")
	
	// Partial Struct Initialization
	personC := Person{
		Name: "Partial",
		Email: "partial@example.com",
	}
	fmt.Println(personC)

	OutputName(personC)
	structAsPointer(&personA)

}


func OutputName(person Person) {
	fmt.Println(person)
}

func structAsPointer(person *Person) {
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Email)
	fmt.Println(person.CreatedAt)
}

