package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
 *	Only functions that start with a capital letter are exported.
 */
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("Error reading file.")
	}

	value, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("Error converting balance.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644)
}
