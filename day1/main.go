package main

import (
	"fmt"
)

func main() {
	// Get data from file & parse it
	rows, err := parserFileToSlice("data/input.txt")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Main function
	count := measurement(rows)

	// Display result
	fmt.Println("Result:", count)
}

// measurement returns int
func measurement(list []int) (count int) {
	for i := 0; i < len(list); i++ {
		if i == 0 {
			continue
		}

		if list[i] > list[i-1] {
			count++
		}
	}

	return count
}
