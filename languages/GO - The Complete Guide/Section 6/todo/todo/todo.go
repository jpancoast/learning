package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

/*
 * The Backticks are struct tags, which can be used to assign metadata
 * The tags in this case are specific to the json package
 *
 * NOTE: for these tags to work, the package / function must be written to use them
 * 	I wonder if it goes the other way as well, ie, reading in a json file
 */
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
