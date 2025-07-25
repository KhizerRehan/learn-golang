package main // main package is the entry point of the program

import (
	"fmt" // fmt package is used to print the output
	"learn-golang/go-with-maximilian/arrayslicesmap"
	"learn-golang/go-with-maximilian/concurrency"
	"learn-golang/go-with-maximilian/fileops"
	"learn-golang/go-with-maximilian/functions"
	"learn-golang/go-with-maximilian/generics"
	"learn-golang/go-with-maximilian/interfaces"
	"learn-golang/go-with-maximilian/pointers"
	"learn-golang/go-with-maximilian/random"
	"learn-golang/go-with-maximilian/structs"
	"learn-golang/go-with-maximilian/user"
)

func main() {
	fmt.Println("Hello World")

	// fileopsWrapper()
	// pointersWrapper()
	// structsWrapper()
	// userWrapper()
	// customTypesWrapper()
	interfacesWrapper()
	// callEmbeddedInterface()
	// callAcceptAnyType()
	// callGenerics()
	// callArraySlicesMap()
	// callArray()	
	// callSlices()
	// callExerceis()
	// callMaps()
	// callFunctions()
	// callAnnonymousFunctions()
	// callClosure()
	// callUnderstandReciever()
	// callVariadicFunction()
	// callConcurrency()
	// callConcurrency1()
}

func fileopsWrapper() {
	fmt.Println("--------------------------------")
	fmt.Println("File Operations")
	fmt.Println("--------------------------------")
		
	content, _ := fileops.ReadFile("./file.txt")
	fmt.Println("File content:")
	fmt.Println(content)
	
}

func pointersWrapper() {
	fmt.Println("--------------------------------")
	fmt.Println("Pointers")
	fmt.Println("--------------------------------")

	fmt.Println("--------------------------------")
	pointers.EmptyPointer()
	fmt.Println("--------------------------------")
	pointers.BasicPointer()
	fmt.Println("--------------------------------")
	pointers.PointerByValue()
	fmt.Println("--------------------------------")
	pointers.PointerByReference()
	fmt.Println("--------------------------------")
	age := 10
	pointers.PointerWithFunction(&age)
	fmt.Println("--------------------------------")
}

func structsWrapper() {
	fmt.Println("--------------------------------")
	fmt.Println("Structs")
	fmt.Println("--------------------------------")
	structs.BasicStruct()

	fmt.Println("--------------------------------")
	fmt.Println("Struct with Receiver")
	fmt.Println("--------------------------------")
	structs.StructWithReceiver()

	fmt.Println("--------------------------------")
	fmt.Println("Struct with By Value or By Reference")
	fmt.Println("--------------------------------")
	structs.StructWithByValueOrByReference()

	fmt.Println("--------------------------------")
	fmt.Println("Struct with Constructor")
	fmt.Println("--------------------------------")
	structs.StructWithConstructor()

	fmt.Println("--------------------------------")
	fmt.Println("Struct with Constructor with Validation")
	fmt.Println("--------------------------------")
	structs.StructWithConstructorWithValidation()

	fmt.Println("--------------------------------")
	fmt.Println("Struct with Pointer")
	fmt.Println("--------------------------------")
	structs.StructWithConstructorWithValidationWithPointer()
}


func userWrapper() {
	fmt.Println("--------------------------------")
	fmt.Println("User")
	fmt.Println("--------------------------------")
	user.GetUserFromExportedStruct()
	fmt.Println("--------------------------------")
	fmt.Println("Admin")
	fmt.Println("--------------------------------")
	user.GetAdminFromExportedStruct()
	fmt.Println("--------------------------------")
	fmt.Println("Static Admin")
	fmt.Println("--------------------------------")
	user.GetStaticAdminFromExportedStruct()
}

func customTypesWrapper() {
	fmt.Println("--------------------------------")
	fmt.Println("Custom Types")
	fmt.Println("--------------------------------")
	structs.CustomTypeWrapper()
}

func interfacesWrapper() {
	fmt.Println("--------------------------------")
	fmt.Println("Interfaces")
	fmt.Println("--------------------------------")
	interfaces.CallSaveTodo()
	interfaces.CallSaveNote()
}

func callEmbeddedInterface() {
	fmt.Println("--------------------------------")
	fmt.Println("Embedded Interfaces")
	fmt.Println("--------------------------------")
	interfaces.CallEmbeddedInterface()
}

func callAcceptAnyType() {
	fmt.Println("--------------------------------")
	fmt.Println("Accept Any Type")
	fmt.Println("--------------------------------")
	interfaces.CallAcceptAnyType()
}

func callGenerics() {
	fmt.Println("--------------------------------")
	fmt.Println("Generics")
	fmt.Println("--------------------------------")
	generics.CallReturnAnyTypeGeneric()
}

func callArraySlicesMap() {
	fmt.Println("--------------------------------")
	fmt.Println("Array Slices Map")
	fmt.Println("--------------------------------")
	arrayslicesmap.CallArraySlicesMap()
}
func callArray() {
	fmt.Println("--------------------------------")
	fmt.Println("Array")
	fmt.Println("--------------------------------")
	arrayslicesmap.CallArray()
}

func callSlices() {
	fmt.Println("--------------------------------")
	fmt.Println("Slices")
	fmt.Println("--------------------------------")
	arrayslicesmap.CallSlices()
}

func callExerceis() {
	fmt.Println("--------------------------------")
	fmt.Println("Exerceis")
	fmt.Println("--------------------------------")
	arrayslicesmap.CallExerceis()
}

func callMaps() {
	fmt.Println("--------------------------------")
	fmt.Println("Maps")
	fmt.Println("--------------------------------")
	arrayslicesmap.LearnMaps()
	arrayslicesmap.MapUsingMake()
	arrayslicesmap.MapUsingType()	
}

func callFunctions() {
	fmt.Println("--------------------------------")
	fmt.Println("Functions")
	fmt.Println("--------------------------------")
	functions.CallFunctions()
}

func callAnnonymousFunctions() {
	fmt.Println("--------------------------------")
	fmt.Println("Annonymous Functions")
	fmt.Println("--------------------------------")
	functions.CallAnnonymousFunctions()
}
func callClosure() {
	fmt.Println("--------------------------------")
	fmt.Println("Closure")
	fmt.Println("--------------------------------")
	functions.CallClosure()
}

func callUnderstandReciever() {
	fmt.Println("--------------------------------")
	fmt.Println("Understand Reciever")
	fmt.Println("--------------------------------")
	structs.CounterWrapper()
}

func callVariadicFunction() {
	fmt.Println("--------------------------------")
	fmt.Println("Variadic Function")
	fmt.Println("--------------------------------")
	random.WrapperForVariadicFunction()
}

func callConcurrency() {
	fmt.Println("--------------------------------")
	fmt.Println("Concurrency")
	fmt.Println("--------------------------------")

	// ------------------------------------------------------------
	// With Done Channel + Wait for Both Channels to Complete
	// ------------------------------------------------------------
	doneChan := make(chan bool)
	concurrency.RunTasksInParallel(doneChan)

	<-doneChan
	fmt.Println("RunTasksInParallel: Tasks are done")


	doneChan2 := make(chan bool)
	concurrency.RunTasksInParallelWithMultipleChannels(doneChan, doneChan2)

	val1 := <-doneChan
	val2 := <-doneChan2
	fmt.Println("val1", val1)
	fmt.Println("val2", val2)
	fmt.Println("RunTasksInParallelWithMultipleChannels: Tasks are done")


    // ------------------------------------------------------------
	// Slice of Bool Channels
	// ------------------------------------------------------------
	// Instead of using multiple channels we can slice the channels
	//  BOTH ABOVE AND BELOW APPROACHES ARE SAME 

	dones := make([]chan bool, 2)
	dones[0] = make(chan bool)
	dones[1] = make(chan bool)

	concurrency.RunTasksInParallelWithMultipleChannels(dones[0], dones[1])

	// Loop through the slice of channels and wait for all of them to complete
	for _,done := range dones {
		<-done
	}
	fmt.Println("RunTasksInParallelWithMultipleChannels: Tasks are done")


	// ------------------------------------------------------------
	// With Wait Group
	// ------------------------------------------------------------
	
	concurrency.RunTasksInParallelWithWaitGroup()


	// ------------------------------------------------------------
	// With Error Channel
	// ------------------------------------------------------------

	errChan := make(chan error)
	concurrency.RunTaskInParallelWithErrorChannel(doneChan, errChan)

	select {
	case <-doneChan:
		fmt.Println("Successfully Done")
	case <-errChan:
		fmt.Println("Error Occured")
	}


}

func callConcurrency1() {

	// ------------------------------------------------------------
	// Slice of Channels
	// ------------------------------------------------------------

	// Create a slice of 3 channels of type interface{}
	
	doneChan3 := make([]chan concurrency.Result, 2)
	// Iterate over the indices of the slice and create a new channel for each index
	// The range keyword returns both the index and the value, but we only need the index here

	for i := range doneChan3 {
		doneChan3[i] = make(chan concurrency.Result)
	}
	
	concurrency.RunTaskWithErrorChannel(doneChan3)
	
	for _, done := range doneChan3 {
		res := <-done
		if res.Error != nil {
			fmt.Println("Error Occurred:", res.Error)
		} else {
			fmt.Println("Successfully Done:", res.Data)
		}
	}

	
}