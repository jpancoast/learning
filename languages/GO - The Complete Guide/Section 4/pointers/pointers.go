package main

func main() {
	age := 32 //This is a regular value

	agePointer := &age
	println("Age Pointer:", agePointer)
	println("Age Pointer Value:", *agePointer)
	/*
	 * Another way to do it:
	 * 	var agePointer *int = &age
	 * OR
	 * 	var agePointer *int
	 * 	agePointer = &age
	 */

	println("Age:", age)
	println("Adult age:", getAdultYears(agePointer))

	println("Pointer shit")
	getAdultYearsPointer(agePointer)
	println("Pointer Age:", age)
}

/*
 * In Go, you can't perform operations on pointers
 */
func getAdultYears(age *int) int {
	return *age - 18
}

func getAdultYearsPointer(age *int) {
	*age -= 18
}
