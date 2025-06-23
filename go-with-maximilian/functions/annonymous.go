// functions return a function
// function wants a function as a parameter value
// We can use annonymous function to pass a function as a parameter value JIT when you need it ather than creating a function and passing it as a parameter value creating in advance

// E,g in functions.go we have a function e.g double, triple, tetraple but not be needed and required only once so maybe defining a function might not be useful

package functions

import "fmt"

func CallAnnonymousFunctions() {
	fmt.Println("--------------------------------")
	fmt.Println("Call Annonymous Functions")
	fmt.Println("--------------------------------")

	numbers := []int{1, 2, 3, 4, 5}

	
	// Annonymous function is a function that is defined without a name

	transformedNumbers := transformAnonymous(&numbers, func(number int) int {return number * 2})
	fmt.Println(transformedNumbers)
}


func transformAnonymous(numbers *[]int, fn func(int) int) []int {
	transformed := []int{}
	for _, value := range *numbers {
		transformed = append(transformed, fn(value))
	}
	return transformed
}

