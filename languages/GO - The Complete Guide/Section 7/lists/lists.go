package main

import (
	"fmt"
)

func main() {
	//Dynamic array
	prices := []float64{2.99, 4.99} //This is technically an empty slice, we didn't give it a length in the []'s
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	prices = append(prices, 5.99, 7.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 292.12}
	prices = append(prices, discountPrices...) //the 3 ...'s basically turhs 'discountPrices' into a comma separated strings
	fmt.Println(prices)
}

// func main() {
// 	prices := [4]float64{1.1, 2.2, 3.3, 4.4}
// 	fmt.Println(prices)

// 	var productNames [7]string //array of length 7
// 	productNames[2] = "Desktopaaaa"
// 	productNames[4] = "Laptop"
// 	productNames[6] = "Desktop"
// 	productNames[5] = "Desktopasdf"
// 	fmt.Println(productNames)

// 	//Slice bitches!  First number is inclusive, second number is exclusive
// 	fmt.Println(prices[1:3])
// 	fmt.Println(productNames[4:7])
// 	fmt.Println(productNames[2:])

// 	/*
// 	 * A SLICE is similar to a pointer within an array, demonstration:
// 	 *	A slice is NOT a copy, it's a reference to a part of that array
// 	 */
// 	fmt.Println(prices)
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 6.67
// 	fmt.Println(prices)

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(len(highlightedPrices)) //len is number of items in an array
// 	fmt.Println(cap(highlightedPrices)) //cap is capacity
// }
