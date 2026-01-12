package d02

import (
	"fmt"
	"os"
	"strings"
)

type Present struct {
	L, W, H int
}

func (p Present) surfaceArea() int {
	return 2 * (p.L*p.W + p.L*p.H + p.W*p.H)
}

func (p Present) ribbonLength() int {
	a, b := p.smallestTwoEdges()
	return 2 * (a + b)
}

func (p Present) smallestTwoEdges() (int, int) {
	if p.L >= p.W && p.L >= p.H {
		return p.W, p.H
	}

	if p.W >= p.L && p.W >= p.H {
		return p.L, p.H
	}
	return p.L, p.W
}
func (p Present) slackArea() int {
	a, b := p.smallestTwoEdges()
	return a * b
}

func (p Present) ribbonBow() int {
	return p.L * p.W * p.H
}
func parseLineToPresent(line string) Present {
	// 2x3x4
	var p Present
	if num, err := fmt.Sscanf(line, "%dx%dx%d", &p.L, &p.W, &p.H); num != 3 || err != nil {
		fmt.Printf("error parsing %s, %v", line, err)
		os.Exit(1)
	}

	return p

}

func Solve(input string) {

	var presents []Present
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		presents = append(presents, parseLineToPresent(line))
	}

	wrappingPaperArea := calculateWrappingPaperArea(presents)
	ribbonLength := calculateLengthOfRibbon(presents)

	fmt.Printf("Part 1: area %d\n", wrappingPaperArea)
	fmt.Printf("Part 2: length of ribbon %d\n", ribbonLength)
}

func calculateWrappingPaperArea(presents []Present) int {
	// surface area
	total := 0

	for _, p := range presents {
		total += p.surfaceArea()
		total += p.slackArea()
	}

	// slack (area of the smallest side)

	return total
}

func calculateLengthOfRibbon(presents []Present) int {
	total := 0

	for _, p := range presents {
		total += p.ribbonLength() + p.ribbonBow()
	}

	return total
}
