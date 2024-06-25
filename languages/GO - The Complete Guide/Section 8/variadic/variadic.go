package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 2, 3, 5, 6)

	fmt.Println(sum)

	sum2 := sumup(numbers...)
	fmt.Println(sum2)
}

/*
 *	The three ...'s mean that the function will accept any number of arguments.
 *	The arguments will be stored in a slice called numbers.
 *	The function will return the sum of all the numbers in the slice.
 *
 *	NOTE: You can accept other parameters as well
 */
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
