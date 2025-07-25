package random

import "fmt"

// function without variadic

func sumUp(numbers []int) int {
	sum :=0

	for _, num := range numbers {
		sum += num
	}
	return sum
}

// function with variadic

func sumUpVariadic(numbers ...int) int {
	sum :=0

	for _, num := range numbers {
		sum += num
	}
	return sum
}

func variadicFunctionWithBaseArray(startingNumber int, numbers2 ...int) int {
	sum :=0

	for _, num := range numbers2 {
		sum += num
	}

	return sum + startingNumber
}

func WrapperForVariadicFunction() {
	// defined slice of numbers explicitly
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumUp(numbers)
	fmt.Println(sum)

	// defined numbers directly
	sum2 := sumUpVariadic(1, 2, 3, 4, 5)
	fmt.Println(sum2)

	sum3 := variadicFunctionWithBaseArray(10, numbers...)	// ... is used to pass the slice of numbers to the function
	fmt.Println(sum3)
}