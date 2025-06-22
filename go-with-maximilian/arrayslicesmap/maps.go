package arrayslicesmap

import "fmt"

func LearnMaps() {

	fmt.Println("--------------------------------")
	fmt.Println("Maps of String to String")
	fmt.Println("--------------------------------")
	
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Twitter": "https://www.twitter.com",
		"Reddit": "https://www.reddit.com",
		"Youtube": "https://www.youtube.com",
		"Instagram": "https://www.instagram.com",
		"Linkedin": "https://www.linkedin.com",
		"Github": "https://www.github.com",
		"Stackoverflow": "https://www.stackoverflow.com",
		"Quora": "https://www.quora.com",
	}


	// Print All Websiste in NextLine

	for key, value := range websites {
		fmt.Println(key, value)
	}

	// Priniting Specific Value

	fmt.Println("--------------------------------")
	fmt.Println("Printing Specific Value")
	fmt.Println("--------------------------------")
	fmt.Println("Key: Stackoverflow, Value:",websites["Stackoverflow"])

	fmt.Println("--------------------------------")
	fmt.Println("Maps of Any Type")
	fmt.Println("--------------------------------")

	mapsOfAnyType := map[interface{}]string{
		"name": "Khizer",
		true: "true",
		false: "false",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	for key, value := range mapsOfAnyType {
		fmt.Println(key, value)
	}

	// Extract Specific Value from Map
	fmt.Println("--------------------------------")
	fmt.Println("Extract Specific Value from Map")
	fmt.Println("--------------------------------")

	mapsOfAnyType["name"] = "Khizer Changed" // Map is Mutable
	fmt.Println(mapsOfAnyType["name"])

	// maps are dynamic in size
	fmt.Println("--------------------------------")
	fmt.Println("Maps are dynamic in size")
	fmt.Println("--------------------------------")
	fmt.Println("Before Addition",mapsOfAnyType)
	fmt.Println("Length", len(mapsOfAnyType))
	mapsOfAnyType["newKey"] = "newValue"
	fmt.Println("After Addition",mapsOfAnyType)
	fmt.Println("Length", len(mapsOfAnyType))

	// Delete Value from Map
	fmt.Println("--------------------------------")
	fmt.Println("Delete Value from Map")
	fmt.Println("--------------------------------")

	fmt.Println("Before Delete",mapsOfAnyType)
	delete(mapsOfAnyType, "name")
	fmt.Println("After Delete",mapsOfAnyType)
	
}

func MapUsingMake() {
	fmt.Println("--------------------------------")
	fmt.Println("Map Using Make")
	fmt.Println("--------------------------------")

	courseRatings := make(map[string]int) // Makeing dynamic map with 0 preserved slot
	courseRatings["Go"] = 5
	courseRatings["Python"] = 4
	courseRatings["Java"] = 3
	courseRatings["JavaScript"] = 2
	courseRatings["C++"] = 1

	fmt.Println("--------------------------------")
	fmt.Println("Map Using Make with 0 preserved slot")
	fmt.Println("--------------------------------")

	fmt.Println(courseRatings)


	// Make with length of 10

	fmt.Println("--------------------------------")
	fmt.Println("Map Using Make with 10 preserved slot")
	fmt.Println("--------------------------------")


	//

	languages := make(map[string]string, 10)
	languages["Go"] = "Go is a programming language"
	languages["Python"] = "Python is a programming language"
	languages["Java"] = "Java is a programming language"
	languages["C++"] = "C++ is a programming language"
	languages["C"] = "C is a programming language"
	languages["C#"] = "C# is a programming language"
	languages["C++"] = "C++ is a programming language"
	languages["C++"] = "C++ is a programming language"

	fmt.Println(languages)
	
}

type floatMap map[string]float64


func MapUsingType() {

	fmt.Println("--------------------------------")
	fmt.Println("Map Using Type")
	fmt.Println("--------------------------------")


	courseRatings := floatMap{
		"Go": 5.0,
		"Python": 4.0,
		"Java": 3.0,
		"JavaScript": 2.0,
		"C++": 1.0,
	}

	courseRatings.print()

	fmt.Println("--------------------------------")
	fmt.Println("Map Using Custom Type")
	fmt.Println("--------------------------------")

	languages := make(floatMap, 10)
	languages["Go"] = 5.0
	languages["Python"] = 4.0
	languages["Java"] = 3.0
	languages["JavaScript"] = 2.0
	languages["C++"] = 1.0
	languages.print()	
	languages.printWithKeyAndValue()	

}

func (m floatMap) print() {
	fmt.Println(m)
}

func (m floatMap) printWithKeyAndValue() {
	for key, value := range m {
		fmt.Println(key, value)
	}
}



