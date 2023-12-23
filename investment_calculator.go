package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount, years, expectedReturnRate = 1000.0, 10.0, 5.5
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	// fmt.Printf("Future Value: %.1f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for Inflation): %.1f", futureRealValue)
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
}
