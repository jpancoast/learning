package main

import (
	"fmt"
)

func main() {
	ebt := 100.12

	fmt.Printf("EBT: %v\n", ebt)

	fv := fmt.Sprintf("EBT_FV: %.2f\n", ebt)
	fmt.Print(fv)

	// Use the backtick when doing multiple lines.
	fmt.Printf(`
Line 1
Line 2
	`)

	tempString := `
one
two
	`

	fmt.Printf(tempString)
}
