package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	y2015d01 "github.com/tvdotdev/advent-of-code/y2015/d01"
	y2015d02 "github.com/tvdotdev/advent-of-code/y2015/d02"
	y2015d03 "github.com/tvdotdev/advent-of-code/y2015/d03"
	y2015d04 "github.com/tvdotdev/advent-of-code/y2015/d04"
	y2015d05 "github.com/tvdotdev/advent-of-code/y2015/d05"
)

func main() {
	year := flag.Int("year", 0, "Year of the puzzle")
	day := flag.Int("day", 0, "day of the puzzle")

	flag.Parse()

	dayPath := fmt.Sprintf("y%d/d%02d", *year, *day)

	solutionPath := filepath.Join(dayPath, "solution.go")

	if _, err := os.Stat(solutionPath); os.IsNotExist(err) {
		fmt.Printf("Solutions for year %d, day %d not found on %s", *year, *day, solutionPath)
		os.Exit(1)
	}

	input := ""
	inputPath := filepath.Join(dayPath, "input.txt")
	if dataBuffer, err := os.ReadFile(inputPath); err == nil {
		input = string(dataBuffer)
	}

	switch {

	case *year == 2015 && *day == 1:
		y2015d01.Solve(input)
	case *year == 2015 && *day == 2:
		y2015d02.Solve(input)
	case *year == 2015 && *day == 3:
		y2015d03.Solve(input)

	case *year == 2015 && *day == 4:
		y2015d04.Solve(input)
	case *year == 2015 && *day == 5:
		y2015d05.Solve(input)

	default:
		fmt.Printf("No solutions implented yet for year %d, day %d", *year, *day)
		os.Exit(1)
	}
}
