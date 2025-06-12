package pointers

import (
	"fmt" // fmt package is used to print the output
)

func EmptyPointer() {
	var emptyPointer *int
	fmt.Println("emptyPointer", emptyPointer)

	fullPointer := new(int)
	fmt.Println("fullPointer", fullPointer)
	fmt.Println("*fullPointer", *fullPointer) // 0

	*fullPointer = 10
	fmt.Println("*fullPointer", *fullPointer) // 10

	var strPtr *string
	fmt.Println("strPtr", strPtr)

	strPtr = new(string)
	fmt.Println("strPtr", strPtr)
	fmt.Println("*strPtr", *strPtr)

	*strPtr = "Hello"
	fmt.Println("*strPtr", *strPtr)
}

func BasicPointer() {
	age := 20
	agePointer := &age


	fmt.Println("age", age)                 // 20
	fmt.Println("agePointer", agePointer)   // pointer to the address of the variable
	fmt.Println("*agePointer", *agePointer) // dereference the pointer to get the value
}

func PointerByValue() {

	x := 10
	y := 10

	fmt.Println("x Before", x)
	fmt.Println("y Before", y)
	result := addPointerByValue(x, y)
	fmt.Println("x After", x)
	fmt.Println("y After", y)

	fmt.Println("addPointerByValue", result)
}

func addPointerByValue(x, y int) int {
	return x + y
}

func PointerByReference() {
	x := 10
	y := 10

	fmt.Println("x Before", x)
	fmt.Println("y Before", y)
	result := addPointerByReference(&x, &y)
	fmt.Println("x After", x)
	fmt.Println("y After", y)

	fmt.Println("addPointerByReference", result)
}

func addPointerByReference(x, y *int) int {
	*x = *x + 10 // mutate the value of x
	*y = *y + 10 // mutate the value of y
	return *x + *y
}


func PointerWithFunction(age *int) int {
	fmt.Println("age Before", *age)
	*age = *age + 10
	fmt.Println("age After", *age)
	return *age
}