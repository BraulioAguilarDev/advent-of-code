package main

import (
	"bufio"
	"os"
	"strconv"
)

// parserFileToSlice parse file to slice
func parserFileToSlice(path string) ([]int, error) {

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
