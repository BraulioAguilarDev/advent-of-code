package main

import (
	"advent-of-code/parser"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	rows, err := parser.GetStrings("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := Dive(rows)

	fmt.Println("Result: ", result)
}

// Dive function returns an int
func Dive(rows []string) int {
	aim := 0
	total := 0
	h := 0

	for i := 0; i < len(rows); i++ {
		commands := strings.Split(rows[i], " ")
		number, _ := strconv.Atoi(commands[1])

		switch commands[0] {
		case "forward":
			h = h + number
			total = total + (aim * number)
		case "down":
			// v = v + number
			aim = aim + number
		case "up":
			// v = v - number
			aim = aim - number
		}
	}

	return h * total
}
