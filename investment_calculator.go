package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// var investmentAmount, years, expectedReturnRate = 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	// fmt.Printf("Future Value: %.1f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for Inflation): %.1f", futureRealValue)
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
