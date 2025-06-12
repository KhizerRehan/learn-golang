package structs

import (
	"errors"
	"fmt"
)

type Animal struct {
	Name string
	Sound string
}


func StructWithReceiver() {

	dog := Animal{
		Name: "Dog",
		Sound: "Woof",
	}

	fmt.Println(dog.OutputName());
	dog.MakeSound()

	cat := Animal{
		Name: "Cat",
		Sound: "Meow",
	}

	fmt.Println(cat.OutputName());
	cat.MakeSound()

}
func StructWithByValueOrByReference() {

	dog := Animal{
		Name: "Dog",
		Sound: "Woof",
	}


	fmt.Println(dog.OutputName());
	dog.MakeSound()

	cat := Animal{
		Name: "Cat",
		Sound: "Meow",
	}


	fmt.Println(cat.OutputName());
	cat.MakeSound()
	fmt.Println("--------------------------------")

	// This will not change the original struct
	fmt.Println("Clear Name By Value")
	dog.clearNameByValue()
	fmt.Println(dog.OutputName());

	fmt.Println("Clear Name By Reference")
	dog.clearNameByReference()
	fmt.Println(dog.OutputName());

}

// --------------------------------
// 	Struct with Receiver
// --------------------------------


// Receiver function is defined which automatically binds to struct level
func (a Animal) MakeSound() {
	fmt.Println(a.Name, "says", a.Sound)
}

func (a Animal) OutputName() string {
	return a.Name + " Sounds" 
}

// This will NOT change the original struct
func (a Animal) clearNameByValue() {
	a.Name = ""
}

// This will CHANGE the original struct
func (a *Animal) clearNameByReference() {
	a.Name = ""
}

// --------------------------------
// 	Struct with Constructor
// --------------------------------

func StructWithConstructor() {
	cat := NewAnimal("Cat", "Meow")
	cat.MakeSound();
}

func StructWithConstructorWithValidation() {
	cat, err := NewAnimalWithValidation("", "Meow")
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cat.MakeSound();
}
	

// Constructor function
func NewAnimal(name string, sound string) (Animal) {
	// This is the way to create a new struct
	// It is not a constructor function
	// It is a function that returns a struct
	// It is a function that returns a struct
	return Animal{
		Name: name,
		Sound: sound,
	}
}


func NewAnimalWithValidation(name string, sound string) (Animal, error) {
	if name == "" || sound == "" {
		return Animal{}, errors.New("Name and Sound are required")
	}

	return Animal{
		Name: name,
		Sound: sound,
	}, nil
}

func StructWithConstructorWithValidationWithPointer() {
	cat, err := NewAnimalPointer("", "wooofy")
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cat.MakeSound();
}


// --------------------------------
// 	Struct with Pointer
// --------------------------------

func NewAnimalPointer(name string, sound string) (*Animal, error) {

	if name == "" || sound == "" {
		return nil, errors.New("Pointer e.g Name and Sound are required")
	}

	// This is the way to create a new struct
	// It is not a constructor function
	// It is a function that returns a struct
	// It is a function that returns a struct
	return &Animal{
		Name: name,
		Sound: sound,
	}, nil
}





	


