package interfaces

import "fmt"

type Saver interface {
	SaveFile() error
}

type Printer interface {
	Print()
}

type SaverPrinter interface {
	Saver	// This is a embedded interface
	Printer	// This is a embedded interface
	CustomDisplayer() // This is a custom method
}

func SaveAndPrint(s SaverPrinter) {
	s.SaveFile()
	s.Print()
	s.CustomDisplayer()
}


// Concreate Example to use embedded interface


type EmbeddedSaverPrinter struct {
	Name string
	Saver
	Printer
}

func (e *EmbeddedSaverPrinter) GetName() string {
	return e.Name
}

func (e *EmbeddedSaverPrinter) SaveFile() error {
	fmt.Println("Saving file...")
	return nil
}

func (e *EmbeddedSaverPrinter) Print() {
	fmt.Println("Printing...")
}

func (e *EmbeddedSaverPrinter) CustomDisplayer() {
	fmt.Println("Custom Displayer...")
}

func NewEmbeddedSaverPrinter(name string) *EmbeddedSaverPrinter {
	return &EmbeddedSaverPrinter{
		Name: name,
	}
}


func CallEmbeddedInterface() {
	embeddedSaverPrinter := NewEmbeddedSaverPrinter("Name: EmbeddedSaverPrinter")

	fmt.Println(embeddedSaverPrinter.GetName())

	embeddedSaverPrinter.Print()	
	embeddedSaverPrinter.SaveFile()
	embeddedSaverPrinter.CustomDisplayer()
}
