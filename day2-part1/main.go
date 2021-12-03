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

	h, v := Dive(rows)

	fmt.Println("Result: ", h*v)
}

// Dive funcs calculates the horizontal and vertical (depth) position
func Dive(rows []string) (h, v int) {
	for i := 0; i < len(rows); i++ {
		commands := strings.Split(rows[i], " ")
		number, _ := strconv.Atoi(commands[1])

		switch commands[0] {
		case "forward":
			h = h + number
		case "down":
			v = v + number
		case "up":
			v = v - number
		}
	}

	fmt.Println("Horizontal:", h)
	fmt.Println("Vertical:", v)

	return h, v
}
