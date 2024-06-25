package main

import "fmt"

func main() {
	//First 'string' is the type of they key. Second string is the type of the value
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Facebook"] = "https://facebook.com"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
}
