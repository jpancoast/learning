package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// var investmentAmount, years float64 = 1000, 10
	var investmentAmount, years float64
	expectedReturnRate := 5.5 //Use := when type should be inferred

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) //& makes it a pointer

	fmt.Print("Number of Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	/*
		futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
		futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	*/
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	altFutureValue, altFutureRealValue := altCalculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	fmt.Println()

	fmt.Println(altFutureValue)
	fmt.Println(altFutureRealValue)
}

/*
 * That bit "(float64, float64)" is what the function is expected to return, in this case, to float64's.
 * 	If just returning one value, it can be like "float64", no parens needed
 */
func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

/*
 * An Alternative way
 *	Notice "(fv float64, rfv float64)" and return is empty
 */
func altCalculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return //Returns the varialbes that are created in the function definition
}
