package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Println("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)

}

func outputText(text string, text2 string) {
	fmt.Println(text, text2)
}
