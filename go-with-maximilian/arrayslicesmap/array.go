package arrayslicesmap

import "fmt"

func CallArray() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	arrayOfStrings := [3]string{"Hello", "World", "Go"}
	fmt.Println(arrayOfStrings)

	arrayOfTemperatures := [3]float64{20.5, 21.5, 22.5}
	fmt.Println(arrayOfTemperatures)

	arrayOfBooleans := [3]bool{true, false, true}
	fmt.Println(arrayOfBooleans)

	arrayOfMixedTypes := [3]any{1, "Hello", true}
	fmt.Println(arrayOfMixedTypes)	
	
	arrayOfMixedTypes2 := [3]interface{}{1, "Hello", true}
	fmt.Println(arrayOfMixedTypes2)

	// Using var 

	// Using var
	var arrayOfStrings2 [3]string
	arrayOfStrings2[0] = "Hello"
	arrayOfStrings2[1] = "World"
	arrayOfStrings2[2] = "Go"
	fmt.Println(arrayOfStrings2)

	var arrayOfTemperatures2 [3]float64
	arrayOfTemperatures2[0] = 20.5
	arrayOfTemperatures2[1] = 21.5
	arrayOfTemperatures2[2] = 22.5
	fmt.Println(arrayOfTemperatures2)

	var arrayOfBooleans2 [3]bool
	arrayOfBooleans2[0] = true
	arrayOfBooleans2[1] = false
	arrayOfBooleans2[2] = true
	fmt.Println(arrayOfBooleans2)


	var arrayOfMixedTypes3 [3]any
	arrayOfMixedTypes3[0] = 1
	arrayOfMixedTypes3[1] = "Hello"
	arrayOfMixedTypes3[2] = true
	fmt.Println(arrayOfMixedTypes3)


	// Empty Array
	var emptyArray [3]int
	fmt.Println(emptyArray)

	var emptyArray2 [3]string
	fmt.Println(emptyArray2)

	var emptyArray3 [3]float64
	fmt.Println(emptyArray3)
	
	var emptyArray4 [3]bool
	fmt.Println(emptyArray4)

	var emptyArray5 [3]any
	fmt.Println(emptyArray5)
	
	
}

