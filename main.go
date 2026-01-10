package main

import (
	"fmt"
	"os"
)

func main() {
	dataBuffer, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	instructions := string(dataBuffer)

	// Part 1: final floor
	floor := getFloorByInstructions(instructions)
	fmt.Printf("Part 1 - Final floor: %d\n", floor)

	// Part 2: position when first entering basement
	position := findBasementPosition(instructions)
	fmt.Printf("Part 2 - First basement position: %d\n", position)
}

func getFloorByInstructions(instructions string) int {
	floor := 0

	for _, instruction := range instructions {
		if instruction == '(' {
			floor += 1
		} else if instruction == ')' {
			floor -= 1
		}
	}
	return floor
}

func findBasementPosition(instructions string) int {
	floor := 0

	for index, instruction := range instructions {
		if instruction == '(' {
			floor += 1
		} else if instruction == ')' {
			floor -= 1
		}

		if floor == -1 {
			return index + 1
		}
	}
	return 0
}
