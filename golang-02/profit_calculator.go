package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate, earningsBeforeTax, earningsAfterTax, ratio float64
	fmt.Println("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Println("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio = earningsBeforeTax / profit
	fmt.Println("Earnings before tax is", earningsBeforeTax)
	fmt.Println("Earnings after tax is", earningsAfterTax)
	fmt.Println("Profit is", profit)
	fmt.Println("Ratio is", ratio)
}
