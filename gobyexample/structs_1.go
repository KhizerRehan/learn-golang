package main

import "fmt"



type Toy struct {
	Name string
	Price int
}

func (t Toy) GetNameAndPrice() (string, int) {
	return t.Name, t.Price
}

func (t* Toy) SetNameAndPrice(name string, price int) {
	t.Name = name
	t.Price = price
}


func NewToy(name string, price int) *Toy {
	return &Toy{
		Name: name,
		Price: price,
	}
}


func main() {

	toy1 := NewToy("Car", 10)
	toy2 := NewToy("Ball", 20)

	fmt.Println(toy1.GetNameAndPrice())
	fmt.Println(toy2.GetNameAndPrice())

	fmt.Println("--------------------------------")
	fmt.Println("Edit Toy-1 Name and Price")
	fmt.Println("--------------------------------")

	toy1.SetNameAndPrice("Bike", 30)

	fmt.Println(toy1.GetNameAndPrice())
	fmt.Println(toy2.GetNameAndPrice())

	fmt.Println("--------------------------------")

}


// Methods can be defined for either pointer or value receiver types