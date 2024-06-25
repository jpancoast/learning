package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	/*
	 *	'make' in this case creates a slice 2 empty elements
	 *	5 is the capacity of the slice
	 */
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	//So after these appends, the array has 4 elements

	fmt.Println(userNames)

	/*
	 *	'make' in this case creates a map with 3 empty elements
	 *
	 *	Really what make is for is pre-allocating memory for things so
	 *		that it's not re-allocating memory every time you add a new
	 *		element to the map.
	 */
	courseRatings := make(map[string]float64, 3)
	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 3.5
	fmt.Println(courseRatings)

	/*
	 *	Type Aliases
	 */
	courseRatingsFloatMap := make(floatMap, 3)
	courseRatingsFloatMap["Go"] = 4.5
	courseRatingsFloatMap["Python"] = 3.5
	fmt.Println("Type Aliases")
	courseRatingsFloatMap.output()

	for index, value := range userNames {
		fmt.Println(index, ":", value)
	}

	for k, v := range courseRatings {
		fmt.Println(k, ":", v)
	}
}
