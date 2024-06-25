package main

import (
	"fmt"

	"compnor.local/account-balance/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "Balance.txt"

/*
 *	For 3rd party packages, we need to install / import them
 *	`go get github.com/Pallinder/go-randomdata`
 */
func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR:", err, " Exiting")
		fmt.Print("--------------------\n\n")
		panic("Can't continue without the balance!")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Phone number:", randomdata.PhoneNumber())

	for { //This is an infinite loop
		presentOptions()

		var choice int
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

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
