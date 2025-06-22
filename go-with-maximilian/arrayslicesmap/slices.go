package arrayslicesmap

import "fmt"


func CallSlices() {
	arrOfStrings := []string{"Hello", "World", "Khizer", "Rehan", "Final"}
	fmt.Println(arrOfStrings)

	arrOfTemperatures := []float64{20.5, 21.5, 22.5, 23.5, 24.5, 25.5}
	fmt.Println(arrOfTemperatures)


	slicesOfStrings := arrOfStrings[2:4]
	fmt.Println(slicesOfStrings)

	slicesOfTemperatures := arrOfTemperatures[3:6]
	fmt.Println(slicesOfTemperatures)

	ArratOfFixedSize()
	CallSlices2()
	CallSlices3()
	AppendSlices()
	EmptySlices()
	MakeSlices()
	MakeSlices()
}

func ArratOfFixedSize() {
	
	fmt.Println("--------------------------------")
	fmt.Println("Array of Fixed Size")
	fmt.Println("--------------------------------")

	arrOfStrings := [3]string{"Hello", "World", "Khizer"}
	fmt.Println(arrOfStrings)


	// NOTE: You can't append to an array of fixed size
	// You can only append to a slice
	// So you need to convert the array to a slice before appending

	/*
	sliceOfStrings := arrOfStrings[:]
	sliceOfStrings = append(sliceOfStrings, "Rehan")
	fmt.Println(sliceOfStrings)

	*/
}

func CallSlices2() {

	arrOfStrings := []string{"Hello", "World", "Khizer", "Rehan", "Final"}
	fmt.Println(arrOfStrings)


	fmt.Println("--------------------------------")
	fmt.Println("Omit Starting Index")
	fmt.Println("--------------------------------")

	omitStartingIndex := arrOfStrings[:3]
	fmt.Println(omitStartingIndex)

	fmt.Println("--------------------------------")
	fmt.Println("Omit Ending Index")
	fmt.Println("--------------------------------")

	omitEndingIndex := arrOfStrings[3:]
	fmt.Println(omitEndingIndex)

	fmt.Println("--------------------------------")
	fmt.Println("Omit Starting and Ending Index")
	fmt.Println("--------------------------------")

	omitStartingAndEndingIndex := arrOfStrings[:]
	fmt.Println(omitStartingAndEndingIndex)
	fmt.Println("Length",len(omitStartingAndEndingIndex))
	fmt.Println("Capacity",cap(omitStartingAndEndingIndex))

	fmt.Println("--------------------------------")
	fmt.Println("Create Slice from Another Slice")
	fmt.Println("--------------------------------")


	createSliceFromAnotherSlice := omitStartingAndEndingIndex[0:2]
	fmt.Println(createSliceFromAnotherSlice)
	fmt.Println("Length",len(createSliceFromAnotherSlice))
	fmt.Println("Capacity",cap(createSliceFromAnotherSlice))

	// Resize Slice
	fmt.Println("--------------------------------")
	fmt.Println("Resize Slice")
	fmt.Println("--------------------------------")

	omitStartingAndEndingIndex = omitStartingAndEndingIndex[0:5]
	fmt.Println(omitStartingAndEndingIndex)
	fmt.Println("Length",len(omitStartingAndEndingIndex))
	fmt.Println("Capacity",cap(omitStartingAndEndingIndex))
	

	fmt.Println("--------------------------------")
	fmt.Println("Update Value in Slice")
	fmt.Println("--------------------------------")

	fmt.Println("Before",omitStartingAndEndingIndex)
	createSliceFromAnotherSlice[0] = "HelloChangedToHi"
	fmt.Println("After",omitStartingAndEndingIndex)
}

func CallSlices3() {
	fmt.Println("--------------------------------")
	fmt.Println("CallSlices3")
	fmt.Println("--------------------------------")

	arrOfStrings := []string{"Hello", "World", "Khizer", "Rehan", "Final"}
	fmt.Println(arrOfStrings)
	fmt.Println("Length",len(arrOfStrings))
	fmt.Println("Capacity",cap(arrOfStrings))

	fmt.Println("--------------------------------")
	fmt.Println("Append Slice")
	fmt.Println("--------------------------------")

	appendSlice := append(arrOfStrings, "NewValue1", "NewValue2")
	fmt.Println("Length",len(appendSlice)) 
	fmt.Println("Capacity",cap(appendSlice))

	fmt.Println("After",arrOfStrings) // does not change original array
	fmt.Println("After",appendSlice) // creates new array and appends new value

	// remove value from slice
	fmt.Println("--------------------------------")
	fmt.Println("Remove Value from Slice")
	fmt.Println("--------------------------------")

	// Start from index 0 and remove 1 value from the end
	removeValueFromSlice := appendSlice[:len(appendSlice)-1]
	fmt.Println("After",removeValueFromSlice)

	// Start from index 0 and remove 2 values from the end
	removeValueFromSlice = appendSlice[:len(appendSlice)-2]
	fmt.Println("After",removeValueFromSlice)

	// Start from index 0 and remove 3 values from the end
	removeValueFromSlice = appendSlice[:len(appendSlice)-3]
	fmt.Println("After",removeValueFromSlice)

	// Start from index 2 and remove 3 values from the end
	removeValueFromSlice = appendSlice[2:len(appendSlice)-3]
	fmt.Println("After",removeValueFromSlice)
}

func AppendSlices() {

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	slice3 := []int{7, 8, 9}
	slice4 := []int{10, 11, 12}

	allSlices := append(slice1, slice2...)
	fmt.Println("All Slices",allSlices)

	allSlices = append(allSlices, slice3...)
	fmt.Println("All Slices",allSlices)

	allSlices = append(allSlices, slice4...)
	fmt.Println("All Slices",allSlices)
	
}

//  This is of Video 122 That I covered in my go learning

func EmptySlices() {

	fmt.Println("--------------------------------")
	fmt.Println("Empty Slices")
	fmt.Println("--------------------------------")

	userNames := []string{} ;// create empty slice

	// userNames[0] = "Khizer" // Will throw error because slice is empty try to uncomment it

	userNames = append(userNames, "Khizer")
	userNames = append(userNames, "Rehan")
	userNames = append(userNames, "Javed")

	userNames[0] = "Khizer Changed"
	userNames[1] = "Rehan Changed"
	userNames[2] = "Javed Changed"

	fmt.Println("User Names",userNames)

	fmt.Println("--------------------------------")
	fmt.Println("Empty Slice")
	fmt.Println("--------------------------------")

	emptySlice := []string{}
	fmt.Println("Empty Slice",emptySlice)
}


func MakeSlices() {

	fmt.Println("--------------------------------")
	fmt.Println("Make Slices")
	fmt.Println("--------------------------------")

	makeSlices := make([]int, 3)
	fmt.Println("Make Slices",makeSlices)

	makeSlices = append(makeSlices, 1, 2, 3, 4, 5)
	fmt.Println("Make Slices",makeSlices)

	makeSlices = append(makeSlices, 1, 2, 3, 4, 5)
	fmt.Println("Make Slices",makeSlices)

	// Preserved Slots can be assigned to new value
	makeSlices[0] = 100
	fmt.Println("Make Slices",makeSlices)

	// Preserved Slots can be assigned to new value
	makeSlices[1] = 200
	fmt.Println("Make Slices",makeSlices)

	// Preserved Slots can be assigned to new value
	makeSlices[2] = 300
	fmt.Println("Make Slices",makeSlices)


	// Makw with 3 slots and 10 capacity
	makeSlices = make([]int, 3, 10)
	fmt.Println("Make Slices",makeSlices)

	// Note it will utilize the 10 capacity and will create new array after capacity is full; it is sort of optimal way to use memory
	makeSlices = append(makeSlices, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Make Slices With 10 Capacity",makeSlices)

	

}	
