package arrayslicesmap

import "fmt"


type Product struct {
	ID int
	Name string
	Price float64
}


func GetProducts() []Product {
	return []Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
		{ID: 3, Name: "Product 3", Price: 300},
	}
}


func printProducts(products []Product) {
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
	}
}

func CallArraySlicesMap() {
	products := GetProducts()
	printProducts(products)
}