package d10

import (
	"fmt"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 10, Solve)
}

func Solve(input string) {

	prev := input
	for i := range 40 {
		prev = lookAndSay(prev)
		fmt.Printf("%2d : %d\n", i, len(prev))
	}

	part1 := len(prev)
	fmt.Printf("Part 1: length after 40 iterations %d\n", part1)

	for range 10 {
		prev = lookAndSay(prev)
	}
	part2 := len(prev)
	fmt.Printf("Part 2: legnth after 50 iterations %d", part2)
}

type Pair struct {
	rune byte
	freq int
}

func lookAndSay(seed string) string {
	// 11221

	// [][]
	// []{}
	pairs := make([]Pair, 0)

	// r is byte
	r := seed[0]
	freq := 1

	for i := 1; i < len(seed); i++ {
		if r == seed[i] {
			freq += 1
		} else {
			pairs = append(pairs, Pair{rune: r, freq: freq})
			r = seed[i]
			freq = 1
		}
	}
	// for the last elements group / pair
	pairs = append(pairs, Pair{rune: r, freq: freq})

	// first solution / very slow
	// output := ""
	// for _, pair := range pairs {
	// 	output = fmt.Sprintf("%s%d%c", output, pair.freq, pair.rune)
	// }
	var stringBuilder strings.Builder

	for _, pair := range pairs {
		// converting the freq for example 3 -> unicode in bytes by adding the rune '0'
		//
		stringBuilder.WriteByte(byte('0' + pair.freq))
		stringBuilder.WriteRune(rune(pair.rune))
	}

	return stringBuilder.String()

}
