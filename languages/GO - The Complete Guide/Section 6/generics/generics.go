package main

func main() {
	result := add(1, 2)
	println(result)
}

/*
 * 'T' is a placeholder type
 * This int | float64 | string is called a type constraint
 * It is used to specify the type of the arguments passed to the function
 *		IOW, T can be a n int, float64, or string, but not a float32, for example
 */
func add[T int | float64 | string](a, b T) T {
	return a + b
}
