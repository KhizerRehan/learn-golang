package interfaces

import "fmt"

// accept any type that implements the interface
func AcceptAnyType(value interface{}) {

	// Type assertion to get the type of the interface
	switch v := value.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Int:", v)
	case float64:
		fmt.Println("Float:", v)
	case bool:
		fmt.Println("Bool:", v)
	case []int:
		fmt.Println("Slice of int:", v)
	case map[string]string:
		fmt.Println("Map of string:", v)
	case map[string]int:
		fmt.Println("Map of int:", v)
	case map[string]bool:
		fmt.Println("Map of bool:", v)
	default:
		fmt.Println("Unknown type:", v)
	}
}


// Incase we need to check specific type of the interface
// as typedValue, ok := value.(string)
// we can use this syntax to check the type of the interface
// typedValue -> can't be reassigned and compiler would know 
// it will be of type string
// ok -> is a boolean that tells us if the type assertion was successful

func AcceptAnyTypeAlternateSyntax(value interface{}) {

	typedValue, ok := value.(string)

	if ok {
		fmt.Println("String:", typedValue)
	}

	typedValue1, ok1 := value.(int)

	if ok1 {
		fmt.Println("Int:", typedValue1)
	}

	typedValue2, ok2 := value.(float64)

	if ok2 {
		fmt.Println("Float:", typedValue2)
	}

	typedValue3, ok3 := value.(bool)

	if ok3 {
		fmt.Println("Bool:", typedValue3)
	}

	typedValue4, ok4 := value.([]int)

	if ok4 {
		fmt.Println("Slice of int:", typedValue4)
	}

	typedValue5, ok5 := value.(map[string]string)

	if ok5 {
		fmt.Println("Map of string:", typedValue5)
	}

	typedValue6, ok6 := value.(map[string]int)

	if ok6 {
		fmt.Println("Map of int:", typedValue6)
	}

	typedValue7, ok7 := value.(map[string]bool)

	if ok7 {
		fmt.Println("Map of bool:", typedValue7)
	}
}


func ReturnAnyType(a, b interface{}) interface{} {

	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return fmt.Sprintf("%.2f", aFloat + bFloat)
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}

	aBool, aIsBool := a.(bool)
	bBool, bIsBool := b.(bool)

	if aIsBool || bIsBool {
		return aBool || bBool
	}

	return nil
}


// Example

func  CalculateSumUsingInterface(a, b interface{}) int {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt * bInt
	}
	return 0
}


func CallAcceptAnyType() {
	AcceptAnyType("Hello")
	AcceptAnyType(123)
	AcceptAnyType(123.45)
	AcceptAnyType(true)
	AcceptAnyType([]int{1, 2, 3})
	AcceptAnyType(map[string]int{"one": 1, "two": 2, "three": 3})
	AcceptAnyType(map[string]string{"one": "1", "two": "2", "three": "3"})
	AcceptAnyType(map[string]bool{"one": true, "two": false, "three": true})

	fmt.Println("--------------------------------")
	fmt.Println("Accept Any Type Alternate Syntax")
	fmt.Println("--------------------------------")
	AcceptAnyTypeAlternateSyntax("Hello")
	AcceptAnyTypeAlternateSyntax(123)
	AcceptAnyTypeAlternateSyntax(123.45)
	AcceptAnyTypeAlternateSyntax(true)
	AcceptAnyTypeAlternateSyntax([]int{1, 2, 3})
	AcceptAnyTypeAlternateSyntax(map[string]int{"one": 1, "two": 2, "three": 3})
	AcceptAnyTypeAlternateSyntax(map[string]string{"one": "1", "two": "2", "three": "3"})
	AcceptAnyTypeAlternateSyntax(map[string]bool{"one": true, "two": false, "three": true})

	fmt.Println("--------------------------------")
	fmt.Println("Return Any Type")
	fmt.Println("--------------------------------")
	fmt.Println(ReturnAnyType(1, 2))
	fmt.Println(ReturnAnyType("khizer", "rehan"))
	fmt.Println(ReturnAnyType(1.0, 2.0))
	fmt.Println(ReturnAnyType(true, false))


	fmt.Println("--------------------------------")
	fmt.Println("Calculate Sum Using Interface")
	fmt.Println("--------------------------------")
	fmt.Println(CalculateSumUsingInterface(1, 2)) // valid number
	fmt.Println(CalculateSumUsingInterface(1, "2")) // invalid number
	
}
