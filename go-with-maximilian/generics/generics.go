package generics

import "fmt"

// This is a ANTI-GENERIC CODE
// I have Copied the code from special_interface.go file


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


//  With Generics
func ReturnAnyTypeGeneric[T int | float64 | string](a, b T) T {
	return a + b
}


func CallReturnAnyTypeGeneric() {
	fmt.Println(ReturnAnyTypeGeneric(1, 2))
	fmt.Println(ReturnAnyTypeGeneric(1.0, 2.0))
	fmt.Println(ReturnAnyTypeGeneric("khizer", "rehan"))
}


