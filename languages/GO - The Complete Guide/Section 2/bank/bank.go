package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "Balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("Error reading balance file.")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("Error converting balance.")
	}

	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR:", err, " Exiting")
		fmt.Print("--------------------\n\n")
		//panic("Can't continue without the balance!")
	}

	fmt.Println("Welcome to Go Bank")

	for { //This is an infinite loop
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		//	fmt.Print("\nYour Choice: ")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		fmt.Println("\nYour Choice was:", choice)

		//wantsCheckBalance := choice == 1  //This works, and sets true or false if choice is 1

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		case 2:
			fmt.Print("Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid, must be > 0")
				//return //This is being used like a break it looks like
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated:", accountBalance)

			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("How much do you want to withdraw?: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Print("Invalid amount, must withdraw > 0")
				//return
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid Ammount, can't withdraw more than you have")
				//return
				continue
			}

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated:", accountBalance)
		default:
			fmt.Println("Exiting")
			//How to exit explicitly?
			//return //Is this the best way to do it?
			//break	//This is for loops

			fmt.Println("Thnks for choosing us")
			return
		}
	}
}
