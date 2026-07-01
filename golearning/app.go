package main

import (
	"fmt"
	"golearning/random/concurrency"
)

func main() {
	
	// interfaces.CallInterfaceV1()
	// interfaces.CallInterfaceV3()
	// interfaces.CallInterfaceV2()
	// random.CallArrayAndIterate()
	// random.CallChannelCollector()

	// concurrency.CallConcurrency()
	// concurrency.CallConcurrencyWithGoRoutine()
	concurrency.CallConcurrencyWithBothGoRoutine()



}


func basicRandom(){
	 // define a boola nd rpinv value and type

	 var value bool
	 fmt.Printf("Type: %T, Value: %v\n", value, value)
	
}