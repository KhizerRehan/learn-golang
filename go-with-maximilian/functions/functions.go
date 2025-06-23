package functions

import "fmt"

func CallFunctions() {
	fmt.Println("--------------------------------")
	fmt.Println("Call Functions")
	fmt.Println("--------------------------------")

	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := doubelNumber(&numbers)
	fmt.Println(dNumbers)

	fmt.Println(transform(&numbers, double))
	fmt.Println(transform(&numbers, triple))	
	fmt.Println(transformWithType(&numbers, tetraple))



	fmt.Println("--------------------------------")
	fmt.Println("Function that returns a function / passing return function as a parameter")
	fmt.Println("--------------------------------")


	numbers2 := []int{1, 2, 3, 4, 5}
	fmt.Println(getTransformFn(2)(10))
	fmt.Println(transformWithType(&numbers2, getTransformFn(2)))
	fmt.Println(transformWithType(&numbers2, getTransformFn(3)))
	fmt.Println(transformWithType(&numbers2, getTransformFn(4)))



	fmt.Println("--------------------------------")
	fmt.Println("Function that returns a function")
	fmt.Println("--------------------------------")

	multiplier := getTransformFn(2) // gets double function
	fmt.Println(multiplier(10)) // calls double function
	multiplier = getTransformFn(3) // gets triple function
	fmt.Println(multiplier(10)) // calls triple function
	multiplier = getTransformFn(4) // gets tetraple function
	fmt.Println(multiplier(10)) // calls tetraple function

}

func doubelNumber(numbers *[]int) []int {
	dNumber := []int{}
	for _, value := range *numbers {
		dNumber = append(dNumber, double(value))
	}
	return dNumber
}


// This is a type alias for a function that takes an int and returns an int
type transformFn func(int) int

// Passing function as a parameter value for other functions
func transform(numbers *[]int, fn func(int) int) []int {
	transformed := []int{}
	for _, value := range *numbers {
		transformed = append(transformed, fn(value))
	}
	return transformed
}


func transformWithType(numbers *[]int, fn transformFn) []int {
	transformed := []int{}
	for _, value := range *numbers {
		transformed = append(transformed, fn(value))
	}
	return transformed
}

// function that returns a function

func getTransformFn(multiplier int) transformFn {
	switch multiplier {
	case 2:
		return double
	case 3:
		return triple
	case 4:
		return tetraple
	default:
		return nil
	}
}


func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func tetraple(number int) int {
	return number * 4
}