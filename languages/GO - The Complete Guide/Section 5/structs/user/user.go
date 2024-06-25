package user

import (
	"fmt"
	"time"
)

// upper or lower case is like with functions in packages, if lower case, it's "private", if upper, it's "public"
type User struct {
	firstName string
	lastName  string
	birthdate string

	createdAt time.Time //This is a struct from the time package
}

// this is embedding another struct
type Admin struct {
	email    string
	password string
	User     //This is called "anonymous" embedding
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

/*
 *	You can export a struct via capitalizing the first letter,
 *		but you don't have to export all the fields in that struct,
 *		If something is lower case then it's not exported.
 */
/*
 *	Notice we don't have to de-ref here?
 *		This is a GO shortcut, you COULD do *u.firstName, etc.
 *
 *	This is a method, it's a function that is attached to a struct,
 *		in this case it's attached to the user struct
 *
 * the (u user) part is called a "Receiver"
 *	Receivers are like values passed to functions, either by copy or by reference
 */
func (u User) OutputUserDetailsStructMethod() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

/*
 * Constructor function
 * Why are constructor functions good?  We can have it validate!
 */
func New(firstName, lastName, birthdate string) User {
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

/*
 * Constructor Function with Pointers
 */
func NewPointer(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, fmt.Errorf("Invalid user data")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func OutputUserDetails(u *User) {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}
