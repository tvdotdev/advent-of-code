package d11

import (
	"fmt"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 11, Solve)
}

func Solve(input string) {
	nextPassword := getNextPassword(input)
	fmt.Printf("Part1: next password after [%s] is [%s]\n", input, nextPassword)
	nextNextPassword := getNextPassword(nextPassword)
	fmt.Printf("Part2: next next password after [%s] is [%s]", nextPassword, nextNextPassword)
}

func getNextPassword(old string) string {

	prev := old
	i := 0
	for {
		prev = increment(prev)

		if !hasConfusingCharacters(prev) && hasIncreasingStraightLikeABC_XYZ(prev) && hasTwoNonOverlapingPairs(prev) {
			return prev
		}

		i++

		if i > 10_000_000 {
			break
		}
	}

	return "NO-PASSWORD-AFTER-1M-ITERATION"
}
func increment(old string) string {

	b := []byte(old)
	for i := len(b) - 1; i >= 0; i-- {

		// this compare the byte value of old[i] and the byte value of 'z'
		// a -> b, c -> d
		if b[i] < 'z' {
			b[i]++
			return string(b)
		}
		// if we are on xz -> became xa -> next iteration -> ya
		b[i] = 'a'
	}

	return string(b)
}

func hasIncreasingStraightLikeABC_XYZ(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		// if a +1 = b -> correct ab sequence
		// if b +1 == c -> correct bc sequence
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			return true
		}
	}
	return false
}

func hasConfusingCharacters(s string) bool {
	return strings.ContainsAny(s, "iol")
}

func hasTwoNonOverlapingPairs(s string) bool {
	pairs := 0

	// aabb
	// 0 -> s[0] s[1] ? yes -> pairs <1 , i = i+1 = 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs++
			i = i + 1
		}
	}

	return pairs >= 2
}
