package structs

import "fmt"

type CustomType string

func (c CustomType) Print() {
	fmt.Println(c)
}


func CustomTypeWrapper() {
	customType := CustomType("Hello")
	customType.Print()

	var customType2 CustomType = "Maximilian"
	customType2.Print()
	
}