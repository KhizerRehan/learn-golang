package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println()

	// ---------------------------------

	weekDay := time.Now().Weekday()
	fmt.Println("Weekday", weekDay)

	// You can use commas to separate multiple expressions in the same case statement.
	//  We use the optional default case in this example as well.

	switch weekDay {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// ---------------------------------

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// ---------------------------------

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	//  In this example, the variable t will have the type corresponding to its clause.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("str")
	whatAmI(12.33)
	fmt.Println()

	// Specific Type is defined

	storeSwitchFunctionInVariable := func(i int) {
		switch i {
		case 1:
			println("1 is pressed")
		case 2:
			println("2 is pressed")

		default:
			println("wrong key is pressed")
		}
	}

	storeSwitchFunctionInVariable(1)
	storeSwitchFunctionInVariable(2)
	storeSwitchFunctionInVariable(10)
}
