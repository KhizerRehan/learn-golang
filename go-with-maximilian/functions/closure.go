package functions

import "fmt"




func CallClosure() {
	fmt.Println("--------------------------------")
	fmt.Println("Call Closure")
	fmt.Println("--------------------------------")

	fullName := getFullName("John", "Doe")
	fmt.Println(fullName())


	double := createTransform(2)
	fmt.Println(double(10))

	triple := createTransform(3)
	fmt.Println(triple(10))
}

func getFullName(firstName string, lastName string) func() string {
	return func() string {
		return firstName + " " + lastName
	}
}

func createTransform(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}