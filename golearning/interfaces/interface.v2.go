package interfaces

import "fmt"

func CallInterfaceV2() {

	//interace type

	mystery := interface{}("Hello")
	DescribeType(mystery)
	DescribeType(1)
	DescribeType(1.0)
	DescribeType(true)
	DescribeType(complex(1, 2))
	DescribeType(complex(1, 2))

	// Check Type

	isString, ok := mystery.(string)
	if ok {
		fmt.Printf("mystery is a string: %v\n", isString)
	} else {
		fmt.Println("mystery is not a string")
	}

	isInt, ok := mystery.(int)
	if ok {
		fmt.Printf("mystery is an int: %v\n", isInt)
	} else {
		fmt.Println("mystery is not an int")
	}

}

func DescribeType(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}
