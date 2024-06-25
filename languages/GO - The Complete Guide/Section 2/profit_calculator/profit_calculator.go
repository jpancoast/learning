package main

import (
	"fmt"
	// "math"
)

/*
 * Updates:
 *	function to prompt for input
 *	single function to calculate all the shite
 */
func main() {
	var revenue, expenses, taxRate float64

	/*
		fmt.Print("Revenue: ")
		fmt.Scan(&revenue)

		fmt.Print("Expenses: ")
		fmt.Scan(&expenses)

		fmt.Print("Tax Rate: ")
		fmt.Scan(&taxRate)
	*/
	revenue = prompt("Revenue: ")
	expenses = prompt("Expenses: ")
	taxRate = prompt("Tax Rate: ")
	/*
	 * What to calculate?
	 *	EBT (Earnings before tax)
	 *	Profit (earnings before tax)
	 *	ratio (EBT/Profit)
	 */
	/*
		ebt := revenue - expenses
		profit := ebt * (1 - (taxRate / 100))
		ratio := ebt / profit
	*/
	ebt, profit, ratio := calcThings(revenue, expenses, taxRate)

	/*
	 * Output:
	 *	EBT
	 *	Profit
	 *	ratio
	 */
	//fmt.Println("EBT:", ebt)
	fmt.Printf("EBT: %v\nProfit: %v\nRatio: %.2f\n", ebt, profit, ratio)
	//fmt.Println(profit)
	//fmt.Println(ratio)

	fv := fmt.Sprintf("EBT_FV: %.2f\n", ebt)
	fmt.Print(fv)
}

func prompt(prompText string) float64 {
	var userInput float64

	fmt.Print(prompText)
	fmt.Scan(&userInput)

	return userInput
}

func calcThings(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - (taxRate / 100))
	ratio = ebt / profit

	return ebt, profit, ratio
}
