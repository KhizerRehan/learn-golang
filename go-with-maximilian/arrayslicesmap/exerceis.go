package arrayslicesmap

import "fmt"


type ProductExercise struct {
	ID int
	Name string
	Quantity int
}


func CallExerceis() {
	// Task1
	hobbies := []string{"Sports", "Cooking", "Reading"}
	fmt.Println("Hobbies", hobbies)

	// Task2
	fmt.Println("--------------------------------")
	fmt.Println("Task2")
	fmt.Println("--------------------------------")

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1])
	fmt.Println(hobbies[2])

	// slice of hobbies from 2nd element to last element
	fmt.Println(hobbies[1:3])

	first2HobbiesSlices := hobbies[:2]
	fmt.Println(first2HobbiesSlices)

	fmt.Println(cap(first2HobbiesSlices))


	courseGoals := []string{"Learn Go", "Learn React", "Learn Docker"}

	allCourses := append(hobbies, courseGoals...)
	fmt.Println(allCourses)


	products := []ProductExercise{
		{ID: 1, Name: "Product 1", Quantity: 10},
		{ID: 2, Name: "Product 2", Quantity: 20},
		{ID: 3, Name: "Product 3", Quantity: 30},
	}
	
	fmt.Println("--------------------------------")
	fmt.Println("Products List")
	fmt.Println("--------------------------------")

	fmt.Println(products)

	fmt.Println("--------------------------------")
	fmt.Println("Products List After Append")
	fmt.Println("--------------------------------")

	products = append(products, ProductExercise{ID: 4, Name: "Product 4", Quantity: 40})
	fmt.Println(products)


}
