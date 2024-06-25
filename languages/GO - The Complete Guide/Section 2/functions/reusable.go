package main

import "fmt"

func outputText(text, text2 string) {
	fmt.Println(text)
	fmt.Println(text2)

	firstString, secondString := calculateThings()

	fmt.Println(firstString)
	fmt.Println(secondString)
}

func calculateThings() (string, string) {
	return "this is a string", "second string"
}
