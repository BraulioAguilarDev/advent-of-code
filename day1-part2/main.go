package main

import (
	"advent-of-code/parser"
	"fmt"
)

func main() {
	// Get data from file & parse it
	rows, err := parser.GetInts("input.txt")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Main function
	count := summedMeasurements(rows)

	// Display result
	fmt.Println("Result:", count)
}

func summedMeasurements(list []int) (count int) {
	for index := 0; index < len(list); index++ {
		if index < 3 {
			continue
		}

		sumA := sum(list[index-3 : index])
		sumB := sum(list[index-2 : index+1])

		if sumB > sumA {
			count++
		}
	}

	return count
}

func sum(list []int) int {
	result := 0
	for _, value := range list {
		result += value
	}

	return result
}
