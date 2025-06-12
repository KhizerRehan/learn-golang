package main // main package is the entry point of the program

import (
	"fmt" // fmt package is used to print the output
	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Hello 3rd Party Packages")
	randomData()
}

func randomData() {
	 // Print a random silly name
	 fmt.Println(randomdata.SillyName())

	 // Print a male title
	 fmt.Println(randomdata.Title(randomdata.Male))
 
	 // Print a female title
	 fmt.Println(randomdata.Title(randomdata.Female))
 
	 // Print a title with random gender
	 fmt.Println(randomdata.Title(randomdata.RandomGender))
 
	 // Print a male first name
	 fmt.Println(randomdata.FirstName(randomdata.Male))
 
	 // Print a female first name
	 fmt.Println(randomdata.FirstName(randomdata.Female))
 
	 // Print a last name
	 fmt.Println(randomdata.LastName())
 
	 // Print a male name
	 fmt.Println(randomdata.FullName(randomdata.Male))
 
	 // Print a female name
	 fmt.Println(randomdata.FullName(randomdata.Female))

	 // Print a currency using ISO 4217
	 fmt.Println(randomdata.Currency())
 
	 // Print a random string sampled from a list of strings
	 fmt.Println(randomdata.StringSample("my string 1", "my string 2", "my string 3"))
 
	 // Print a valid random IPv4 address
	 fmt.Println(randomdata.IpV4Address())
 
	 // Print a valid random IPv6 address
	 fmt.Println(randomdata.IpV6Address())

}