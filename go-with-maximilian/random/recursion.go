package random

import "fmt"


func RecursionUsingRecursion(n int) int {
	
	if(n == 1) {
		return 1
	}

	return n * RecursionUsingRecursion(n-1)
}



func RecursionUsingLoop(n int) { 
	factorial := 1

	for i :=factorial; i<=n; i++ {
		factorial = factorial * i // 1*1 =1 , 1*2 = 2, 2*3 = 6, 6*4 = 24, 24*5 = 120
	}
	fmt.Println(factorial)
}