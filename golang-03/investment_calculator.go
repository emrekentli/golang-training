package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Println("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)

}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
