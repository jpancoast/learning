package main

import (
	"fmt"
)

func main() {
	fact := factorial2(5)
	fmt.Println(fact)
}

// calculate factorial, this is the loop way
func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

// The recursive way to do it
func factorial2(number int) int {
	//The exit condition
	if number == 0 {
		return 1
	}

	return number * factorial2(number-1)
}
