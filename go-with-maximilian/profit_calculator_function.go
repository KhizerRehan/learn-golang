package main

import "fmt"

func main() {
	// Profit Calculator Via Functions
	profitCalculatorWithFunction();
}


func profitCalculatorWithFunction() {
	revenue, expenses, taxRate := inputValues();
	profitBeforeTax, profitAfterTax, profitRatio := calculateValues(revenue, expenses, taxRate);
	// profitBeforeTax, profitAfterTax, profitRatio := calculateValuesShortSyntax(revenue, expenses, taxRate);
	outputText("Your profit before tax is: " + fmt.Sprintf("%.02f", profitBeforeTax))
	outputText("Your profit after tax is: " + fmt.Sprintf("%.02f", profitAfterTax))
	outputText("Your profit ratio is: " + fmt.Sprintf("%.02f", profitRatio))
}

func outputText(text string) {
	fmt.Println(text)
}

func inputValues() (float64, float64, float64) {
	
	fmt.Println("--------------------------------");
	fmt.Println("Input Values");
	fmt.Println("--------------------------------");

	revenue := getUserInput("Enter your revenue: ");
	expenses := getUserInput("Enter your expenses: ");
	taxRate := getUserInput("Enter your tax rate (as a percentage %): ");

	return revenue, expenses, taxRate;
}

func getUserInput(text string) float64 {
	var input float64;
	fmt.Println(text);
	fmt.Scanln(&input);
	return input;
}


func calculateValues(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	
	fmt.Println("--------------------------------");
	fmt.Println("Calculating Values");
	fmt.Println("--------------------------------");

	profitBeforeTax := revenue - expenses;
	profitAfterTax := profitBeforeTax * (1 - taxRate / 100);
	profitRatio := profitBeforeTax / profitAfterTax;

	return profitBeforeTax, profitAfterTax, profitRatio;
}


// Info: NOT Using Just shown different syntax
// Alternative Signature
// Shorter Syntax for Parameters and Return Values

func calculateValuesShortSyntax(revenue , expenses , taxRate float64) (profitBeforeTax , profitAfterTax , profitRatio float64) {

	fmt.Println("--------------------------------");
	fmt.Println("Shoter Syntax: Calculating Values");
	fmt.Println("--------------------------------");

	profitBeforeTax = revenue - expenses;
	profitAfterTax = profitBeforeTax * (1 - taxRate / 100);
	profitRatio = profitBeforeTax / profitAfterTax;
	return
}



