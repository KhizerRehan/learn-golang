package main

import (
	"fmt"
)

/*
 zeroval and zeroptr. zeroval has an int parameter,
 so arguments will be passed to it by value. zeroval
 will get a copy of ival distinct from the one in the
 calling function.
*/

func zeroval(ival int) {
	ival = 0
}

/*
zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
The *iptr code in the function body then dereferences the pointer from its memory
address to the current value at that address. Assigning a value to a dereferenced
pointer changes the value at the referenced address.
*/
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	fmt.Println("Points Examples")

	i := 10
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// zeroval doesn’t change the i in main, but zeroptr does because
	// it has a reference to the memory address for that variable.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer value:", i)
	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("pointer address:", &i)

	// C++ Concepts
	var value *int = &i
	fmt.Println("pointer value:", value)
	fmt.Println("pointer address:", *value)
}
