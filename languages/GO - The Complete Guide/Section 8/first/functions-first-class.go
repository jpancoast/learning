package main

import "fmt"

type transformFn = func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//dNumbers := doubleNumbers(&numbers)

	dNumbers := transformNumbers(&numbers, double)
	fmt.Println(dNumbers)

	tNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tNumbers)

	moreNumbers := []int{5, 1, 2}
	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers1 := transformNumbers(&numbers, transformFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)
}

/*
 *	Functions are first class values
 */
func transformNumbers(numbers *[]int, transform transformFn) []int {
	newNumbers := []int{}

	for _, val := range *numbers {
		newNumbers = append(newNumbers, transform(val))
	}

	return newNumbers
}

// This returns a function!
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
