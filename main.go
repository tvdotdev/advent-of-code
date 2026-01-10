package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello world")

	dataBuffer, _ := os.ReadFile("./input")

	instructions := string(dataBuffer)
	floor := getFloorByInstructionsPt2(instructions)
	fmt.Printf("\nfloor is %d", floor)
}

func getFloorByInstructions(instructions string) int {
	floor := 0

	for _, instruction := range instructions {
		if instruction == '(' {
			floor += 1
		}

		if instruction == ')' {
			floor -= 1
		}
	}
	return floor
}

func getFloorByInstructionsPt2(instructions string) int {
	floor := 0

	for index, instruction := range instructions {
		if instruction == '(' {
			floor += 1
		}

		if instruction == ')' {
			floor -= 1
		}

		if floor == -1 {
			return index + 1
		}
	}
	return 0
}
