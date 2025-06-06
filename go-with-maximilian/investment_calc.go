package main

import (
	"fmt"
	"math"
)

func main() {

	withBuiltInTypes()
	goInferredTypeBasedOnColon()
	singleLineDeclarationShorterSyntax()
	singleLineDeclarationExplicitTyping()
	fmt.Println("--------------------------------")
	withUserInput()

}

func withBuiltInTypes() {
	var investmentAmount float64 = 1000
	var expectedReturn float64 = 5.5
	var years int = 10
	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println(futureValue)
}

func goInferredTypeBasedOnColon() {
	investmentAmount := 1000.0
	expectedReturn := 5.5
	years := 10
	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println(futureValue)
}

func singleLineDeclarationShorterSyntax() {
	// Go Infer Types
	investmentAmount, expectedReturn, years := 1000.0, 5.5, 10.0
	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Println(futureValue)
}

func singleLineDeclarationExplicitTyping() {
	var investmentAmount, expectedReturn, years float64 = 1000.0, 5.5, 10 // Golang explicit cast int -> float Internally
	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Println(futureValue)
}

func withUserInput() {

	var investmentAmount float64
	var expectedReturn float64
	var years int
	var inflationRate float64 = 2.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scanln(&investmentAmount)

	fmt.Print("Enter expected return: ")
	fmt.Scanln(&expectedReturn)

	fmt.Print("Enter the number of years: ")
	fmt.Scanln(&years)

	fmt.Print("Enter the inflation rate: ")
	fmt.Scanln(&inflationRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println("Future value: ", futureValue)

	var futureValueAfterInflation = futureValue * math.Pow(1+inflationRate/100, float64(years))
	fmt.Println("Future value after inflation: ", futureValueAfterInflation)

}
