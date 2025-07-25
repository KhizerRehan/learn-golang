package random

import "fmt"


func CallArrayAndIterate() {

	intArrays := []int{1, 2, 3, 4, 5}
	stringArrays := []string{"A", "B", "C", "D", "E"}
	boolArrays := []bool{true, false, true, false, true}
	floatArrays := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	complexArrays := []complex128{complex(1, 2), complex(3, 4), complex(5, 6), complex(7, 8), complex(9, 10)}
	mapArrays := []map[string]int{
		{"a": 1, "b": 2, "c": 3},
		{"d": 4, "e": 5, "f": 6},
	}

	for _, value := range intArrays {
		fmt.Println(value)
	}

	for _, value := range stringArrays {
		fmt.Println(value)
	}

	for _, value := range boolArrays {	
		fmt.Println(value)
	}

	for _, value := range floatArrays {
		fmt.Println(value)
	}

	for _, value := range complexArrays {
		fmt.Println(value)
	}


	fmt.Println("--------------------------------")

	// Iterate over map	value
	for _, value := range mapArrays {
		for key, value := range value {
			fmt.Println(key, value)
		}
	}


}