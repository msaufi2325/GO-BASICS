package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount, years, expectedReturnRate = 1000.0, 10.0, 5.5
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}