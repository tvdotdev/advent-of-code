package d06

import (
	"fmt"
	"os"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 6, Solve)
}

type Point struct {
	X, Y int
}

// turn on 0,0 through 999,999 -> Instruction {Action: "turn on", From : {X: 0, Y:0}} ...
type Instruction struct {
	Action string
	From   Point
	To     Point
}

func Solve(input string) {

	instructionStrings := strings.Split(input, "\n")
	instructions := make([]Instruction, len(instructionStrings))
	for _, line := range instructionStrings {
		if line == "" {
			continue
		}
		instructions = append(instructions, parseLineString(line))
	}

	litLights := howManyLightsAreLit(instructions)

	fmt.Printf("Part 1: %d lights are lit\n", litLights)

	totalBrightness := getTotalBrightness(instructions)
	fmt.Printf("Part 2: total brightness is %d", totalBrightness)

}

func howManyLightsAreLit(instructions []Instruction) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		applyInstruction(&grid, instruction)
	}

	return countLitLights(grid)
}

func getTotalBrightness(instructions []Instruction) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		applyInstruction2(&grid, instruction)
	}

	return calculateTheTotalBrightness(grid)
}

func calculateTheTotalBrightness(grid [][]int) int {
	sum := 0

	for x := range 1000 {
		for y := range 1000 {
			sum += grid[x][y]
		}
	}
	return sum
}

func countLitLights(grid [][]int) int {
	// == x = 0 -> x <= 999
	count := 0

	for x := range 1000 {
		for y := range 1000 {
			if grid[x][y] == 1 {
				count++
			}
		}
	}
	return count
}
func applyInstruction(grid *[][]int, instruction Instruction) {
	for x := instruction.From.X; x <= instruction.To.X; x++ {
		for y := instruction.From.Y; y <= instruction.To.Y; y++ {
			switch instruction.Action {
			case "turn off":
				(*grid)[x][y] = 0
			case "turn on":
				(*grid)[x][y] = 1

			case "toggle":
				// if 1 -> 0 ; 0 -> 1
				// xor ^
				(*grid)[x][y] ^= 1
			}
		}
	}
}

func applyInstruction2(grid *[][]int, instruction Instruction) {
	for x := instruction.From.X; x <= instruction.To.X; x++ {
		for y := instruction.From.Y; y <= instruction.To.Y; y++ {
			switch instruction.Action {
			case "turn off":
				if (*grid)[x][y] > 0 {
					(*grid)[x][y] -= 1
				}
			case "turn on":
				(*grid)[x][y] += 1

			case "toggle":
				(*grid)[x][y] += 2
			}
		}
	}
}

func parseLineString(line string) Instruction {
	// turn on 0,0 through 999,999
	// turn off 499,499 through 500,500
	// toggle 499,499 through 500,500
	var instruction Instruction

	// 0,0 through 999,999
	var coordinates string

	if c, found := strings.CutPrefix(line, "turn off"); found {
		instruction.Action = "turn off"
		coordinates = c
	} else if c, found := strings.CutPrefix(line, "turn on"); found {
		instruction.Action = "turn on"
		coordinates = c
	} else if c, found := strings.CutPrefix(line, "toggle"); found {
		instruction.Action = "toggle"
		coordinates = c
	}

	if num, err := fmt.Sscanf(coordinates, "%d,%d through %d,%d", &instruction.From.X, &instruction.From.Y, &instruction.To.X, &instruction.To.Y); num != 4 || err != nil {
		fmt.Printf("Error parsing %s, %v", coordinates, err)
		os.Exit(1)
	}

	return instruction
}
