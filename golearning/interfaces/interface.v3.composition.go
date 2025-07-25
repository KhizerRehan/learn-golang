package interfaces

import (
	"fmt"
)

// Base interfaces
type Sounder interface {
	Sound() string
}

type Mover interface {
	Move() string
}

type Eater interface {
	Eat() string
}

type AnimalBehavior interface {
	Sounder
	Mover
	Eater
}

// Embedded interfaces
type Animal interface {
	AnimalBehavior
	Name() string
	Age() int
	Color() string
}

// Implementations
type Dog struct {
	name string
	age  int
	color string
}

func (d Dog) Sound() string {
	return "Woof"
}

func (d Dog) Move() string {
	return "Dog runs on four legs"
}

func (d Dog) Eat() string {
	return "Dog eats Meat"
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Age() int {
	return d.age
}

func (d Dog) Color() string {
	return d.color
}

// Wrapper that uses embedded interface
type WhoAmI struct {
	Animal
}

func (a WhoAmI) describeAnimal() string {
	return fmt.Sprintf("Animal Name: %s\nAnimal Sound: %s\nAnimal Move: %s\nAnimal Eat: %s\nAnimal Age: %d\nAnimal Color: %s", 
		a.Name(), a.Sound(), a.Move(), a.Eat(), a.Age(), a.Color())
}

func CallInterfaceV3() {

	dog := Dog{name: "Rottweiler", age: 3, color: "Brown"}
	animal := WhoAmI{Animal: dog}
	fmt.Println(animal.describeAnimal())

}