package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"compnor.local/todo/note"
	"compnor.local/todo/todo"
)

// Interfaces make sure that the structs implement the methods
type saver interface {
	Save() error
}

//type displayer interface {
//	display()
//}

/*
 * This is an embedded interface.
 * It means that the structs that implement this interface will also implement the methods of the embedded interface.
 */
type outputabble interface {
	saver
	Display()
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
}

func outputData(data outputabble) error {
	data.Display()
	return saveData(data)
}

// We want to accept ANY kind of value
func printSomething(value interface{}) {
	fmt.Println(value)
}

// Do different things based on the type of the value passed
func printSomethingSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("String")
	}
}

// Will return ok = true if value is of type int
// "any" is an alias for interface{}
func printSomethingValue(value any) {
	typedVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", typedVal)
		return
	}

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}
}

// Interfaces make sure that the structs implement the methods

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	printSomethingSwitch(123.5)

	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		println("Daving the todo succeeded.")

		userNote.Display()
		err = saveData(userNote)

		if err != nil {
			fmt.Println("Saving the note failed.")
			return
		}

		fmt.Println("Saving the note succeeded!")
	*/
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
