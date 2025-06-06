package main

import (
	"fmt"
)

func main() {
	profitCalculator()
}

func profitCalculator() {
	var revenue float64;
	var expenses float64;
	var taxRate float64;

	var profitRatio float64;
	var profitBeforeTax float64


	fmt.Println("Enter your revenue: ");
	fmt.Scanln(&revenue);

	fmt.Println("Enter your expenses: ");
	fmt.Scanln(&expenses);

	fmt.Println("Enter your tax rate (as a percentage %): ");
	fmt.Scanln(&taxRate);

	profitBeforeTax = revenue - expenses;
	profitAfterTax := profitBeforeTax * (1 - taxRate / 100);

	profitRatio = profitBeforeTax / profitAfterTax;

	fmt.Println("Your profit before tax is: ", profitBeforeTax);
	fmt.Printf("Your profit after tax is: %.02f\n", profitAfterTax);
	fmt.Printf("Your profit ratio is: %.02f\n", profitRatio);

	// In case value needs to be formatted and stored in a variable
	fmt.Println("--------------------------------");
	fmt.Println("Formatted values: ");
	fmt.Println("--------------------------------");	
	formattedPBT := fmt.Sprintf("PBT:: %.02f\n", profitBeforeTax);
	formattedPAT := fmt.Sprintf("PAT: %.02f\n", profitAfterTax);
	formattedPR := fmt.Sprintf("PR: %.02f\n", profitRatio);

	fmt.Println(formattedPBT);
	fmt.Println(formattedPAT);
	fmt.Println(formattedPR);

}