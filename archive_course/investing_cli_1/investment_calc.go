package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5

	var investmentAmount, years, expectedReturnRate float64

	//Get user input with fmt.Scan
	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years you want to invest this cash: ")
	fmt.Scan(&years)

	fmt.Print("Finally, enter the expected rate of return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//Taking inflation into account:
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value, adjusted for inflation:", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value Adjusted for Inflation: %.0f", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Printf("Future Value: %.0f\nFuture Value Adjusted for Inflation: %.0f", futureValue, futureRealValue)
	// Formatting values - can use %.0f to round to nearest whole number
	// or use %.2f to round to 2 decimal places, and so on.

}
