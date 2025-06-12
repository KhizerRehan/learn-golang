package main // main package is the entry point of the program

import (
	"fmt" // fmt package is used to print the output
	"learn-golang/go-with-maximilian/fileops"
	"learn-golang/go-with-maximilian/pointers"
	"learn-golang/go-with-maximilian/structs"
	"learn-golang/go-with-maximilian/user"
)

func main() {
	fmt.Println("Hello World")

	// fileopsWrapper()
	// pointersWrapper()
	// structsWrapper()
	// userWrapper()
	customTypesWrapper()
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