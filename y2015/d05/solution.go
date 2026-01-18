package d05

import (
	"fmt"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 5, Solve)
}

func Solve(input string) {
	santaStrings := strings.Split(input, "\n")

	part1 := solvePart1(santaStrings)

	fmt.Printf("Part 1: %d\n", part1)

	part2 := solvePart2(santaStrings)

	fmt.Printf("Part 2: %d\n", part2)

}

func solvePart1(santaStrings []string) int {
	count := 0
	for _, santaString := range santaStrings {
		if isStringNice(santaString) {
			count++
		}

	}
	return count
}
func solvePart2(santaStrings []string) int {
	count := 0
	for _, santaString := range santaStrings {
		if isStringNice2(santaString) {
			count++
		}

	}
	return count
}

func hasAtLeastThreeVowels(s string) bool {
	count := 0

	for _, r := range s {
		if r == 'a' || r == 'o' || r == 'e' || r == 'i' || r == 'u' {
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

func hasDoubleLetters(s string) bool {

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasForbiddenString(s string) bool {
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}

	for _, forbiddenString := range forbiddenStrings {
		if strings.Contains(s, forbiddenString) {
			return true
		}
	}

	return false
}

func isStringNice(s string) bool {
	return hasAtLeastThreeVowels(s) && hasDoubleLetters(s) && !hasForbiddenString(s)
}

func hasNonOverlappingPari(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		pair := s[i : i+2]
		rest := s[i+2:]
		if strings.Contains(rest, pair) {
			return true
		}
	}
	return false
}

func hasXYXPattern(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		letter1 := s[i]
		letter2 := s[i+2]
		if letter1 == letter2 {
			return true
		}
	}
	return false
}

func isStringNice2(s string) bool {
	return hasNonOverlappingPari(s) && hasXYXPattern(s)
}
