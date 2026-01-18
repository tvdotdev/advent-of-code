package d04

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 4, Solve)
}
func Solve(input string) {
	part1 := findNumberHash(input, "00000")
	fmt.Printf("Part 1: %d", part1)

	part2 := findNumberHash(input, "000000")
	fmt.Printf("Part 2: %d", part2)
}

func findNumberHash(secretKey string, prefix string) int {
	number := 0

	// for number := 1; number <= 10_000_000; number++ { ... }
	// for { }

	for {
		number++
		input := fmt.Sprintf("%s%d", secretKey, number) // as string
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))

		if strings.HasPrefix(hash, prefix) {
			return number
		}

		if number > 10_000_000 {
			fmt.Printf("Error: number reached 10M and no hash")
			return -1
		}
	}
}
