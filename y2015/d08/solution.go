package d08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 8, Solve)
}

func Solve(input string) {
	// lines := strings.Split(input, "\n")
	lines := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })

	part1 := numOfCharactersInCode(lines) - numOfCharactersInMemory(lines)

	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part1/improved: %d\n", solvePart1(lines))

	part2 := numOfDoubleEncoding(lines) - numOfCharactersInCode(lines)
	fmt.Printf("Part 2: %d", part2)

}

func numOfCharactersInCode(lines []string) int {
	count := 0

	for _, line := range lines {
		count += len(line)
	}

	return count
}

func numOfCharactersInMemory(lines []string) int {
	count := 0

	for _, line := range lines {
		parsed, _ := strconv.Unquote(line)
		count += len(parsed)
	}
	return count
}

func numOfDoubleEncoding(lines []string) int {
	count := 0
	for _, line := range lines {
		count += len(strconv.Quote(line))
	}

	return count
}

// improvement

func solvePart1(lines []string) int {
	count := 0

	for _, line := range lines {
		count += len(line)
		parsed, _ := strconv.Unquote(line)
		count -= len(parsed)
	}
	return count
}
