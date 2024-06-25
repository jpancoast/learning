package main

import "fmt"

// 'customString' is an alias to 'string', use cases will be shown later
type customString string

// can attach methods to customString
func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "John"

	name.log()
}
