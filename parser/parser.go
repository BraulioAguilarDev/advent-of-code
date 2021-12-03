package parser

import (
	"bufio"
	"os"
	"strconv"
)

// Refactor functions, are similar -- only change the casting (Atoi)

// GetStrings returns a slice of int from file
func GetInts(path string) ([]int, error) {
	var rows []int

	file, err := os.Open(path)
	if err != nil {
		return rows, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		rows = append(rows, value)
	}

	return rows, nil
}

// GetStrings returns a slice of string from file
func GetStrings(path string) ([]string, error) {
	var rows []string

	file, err := os.Open(path)
	if err != nil {
		return rows, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows, nil
}
