package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Golang Bank")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is $", accountBalance)
		case 2:
			fmt.Println("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Invalid amount")
				break
			}

			fmt.Println("You have deposited $", depositAmount)
			accountBalance += depositAmount
			fmt.Println("Your new balance is $", accountBalance)
		case 3:
			fmt.Println("Enter amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount < 0 {
				fmt.Println("Invalid amount")
				break
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				break
			}

			fmt.Println("You have withdrawn $", withdrawAmount)
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is $", accountBalance)
		case 4:
			fmt.Println("Thank you for banking with us")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
