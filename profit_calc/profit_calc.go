package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	fmt.Println("Welcome to your profit calculator!!")
	fmt.Print("Please enter your expected revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Please enter any expenses you had this year: ")
	fmt.Scan(&expenses)
	fmt.Print("Finally, please enter your tax rate: ")
	fmt.Scan(&taxRate)

	var earningsBeforeTax = revenue - expenses
	var earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)

	var ratio = earningsBeforeTax / earningsAfterTax

	fmt.Println("Your earnings before tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Println("Your earnings after tax: ")
	fmt.Println(earningsAfterTax)

	fmt.Println("Your ratio is: ")
	fmt.Println(ratio)

}
