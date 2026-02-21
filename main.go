package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tvdotdev/advent-of-code/register"
	_ "github.com/tvdotdev/advent-of-code/y2015/d01"
	_ "github.com/tvdotdev/advent-of-code/y2015/d02"
	_ "github.com/tvdotdev/advent-of-code/y2015/d03"
	_ "github.com/tvdotdev/advent-of-code/y2015/d04"
	_ "github.com/tvdotdev/advent-of-code/y2015/d05"
	_ "github.com/tvdotdev/advent-of-code/y2015/d06"
	_ "github.com/tvdotdev/advent-of-code/y2015/d07"
	_ "github.com/tvdotdev/advent-of-code/y2015/d08"
	_ "github.com/tvdotdev/advent-of-code/y2015/d09"
	_ "github.com/tvdotdev/advent-of-code/y2015/d10"
	_ "github.com/tvdotdev/advent-of-code/y2015/d11"
	_ "github.com/tvdotdev/advent-of-code/y2015/d12"
	_ "github.com/tvdotdev/advent-of-code/y2015/d13"
	_ "github.com/tvdotdev/advent-of-code/y2015/d14"
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

	if solve, ok := register.Solutions[*year][*day]; ok {
		solve(input)
	} else {
		fmt.Printf("No solution yet for year %d day %d", *year, *day)
		os.Exit(1)
	}
}
