package main

import (
	"fmt"

	"compnor.local/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser, testUser user

	//This is the struct "literal" notation
	/*
		appUser = user{
			firstName: userFirstName,
			lastName:  userLastName,
			birthdate: userBirthdate,
			createdAt: time.Now(),
		}
	*/

	//This works as well, as long as values are in the correct order
	/*
		testUser = user{
			userFirstName,
			userLastName,
			userBirthdate,
			time.Now(),
		}
	*/

	//Using the constructor function
	//var constructorUser *user.User
	constructorUser, err := user.NewPointer(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("email@email.com", "password")

	//Gathering data from a struct
	// ... do something awesome with that gathered data!

	//fmt.Println(firstName, lastName, birthdate)
	/*
		println("---------------")
		outputUserDetails(&appUser)
		println("---------------")
		outputUserDetails(&testUser)
		println("---------------")
		appUser.outputUserDetailsStructMethod()
		println("---------------")
		appUser.clearUserName()
		appUser.outputUserDetailsStructMethod()
	*/
	println("---------------")
	constructorUser.OutputUserDetailsStructMethod()
	constructorUser.ClearUserName()
	constructorUser.OutputUserDetailsStructMethod()
	println("---------------")
	admin.OutputUserDetailsStructMethod()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
