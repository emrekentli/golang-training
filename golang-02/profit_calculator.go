package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	writeTextAndScan("Enter revenue: ", &revenue)
	writeTextAndScan("Enter expenses: ", &expenses)
	writeTextAndScan("Enter tax rate: ", &taxRate)

	earningsBeforeTax, earningsAfterTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Println("Earnings before tax is", earningsBeforeTax)
	fmt.Println("Earnings after tax is", earningsAfterTax)
	fmt.Println("Profit is", profit)
	fmt.Println("Ratio is", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, earningsAfterTax, profit, ratio
}

func writeTextAndScan(text string, value *float64) {
	fmt.Print(text)
	fmt.Scan(value)
}
