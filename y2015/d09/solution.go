package d09

import (
	"fmt"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 9, Solve)
}

func Solve(input string) {
	distances := make(map[string]map[string]uint)

	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		parse(line, &distances)
	}

	cities := make([]string, 0)
	for city := range distances {
		cities = append(cities, city)
	}

	premutations := getPremutations(cities)

	min, max := getMinMax(premutations, distances)
	fmt.Printf("Part 1: %d\n", min)
	fmt.Printf("Part 2: %d\n", max)

}

func parse(line string, distances *map[string]map[string]uint) {
	// London to Dublin = 123
	from, to, dist := "", "", uint(0)
	if num, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &dist); num != 3 || err != nil {
		panic("something went wrong parsing " + line)
	}

	if (*distances)[from] == nil {
		(*distances)[from] = make(map[string]uint)
	}
	if (*distances)[to] == nil {
		(*distances)[to] = make(map[string]uint)
	}

	// london -> dublin
	(*distances)[from][to] = dist

	// dublin -> london
	(*distances)[to][from] = dist

}

func getPremutations(cities []string) [][]string {
	if len(cities) == 1 {
		return [][]string{cities}
	}

	var results = make([][]string, 0)
	for i, city := range cities {
		remaining := make([]string, len(cities)-1)
		// copy to slice a, slice b
		copy(remaining, cities[:i])
		copy(remaining[i:], cities[i+1:])

		remainingPremutations := getPremutations(remaining)

		// city -> a
		// remainingPremutation [[bc][cb]]
		// [[abc] [acb]]
		newPremutation := make([][]string, 0)
		for _, el := range remainingPremutations {
			newPremutation = append(newPremutation, append([]string{city}, el...))
		}
		results = append(results, newPremutation...)
	}

	return results
}

func getMinMax(premutations [][]string, distances map[string]map[string]uint) (uint, uint) {
	min := ^uint(0) // this will put min = maximum value in uint
	max := uint(0)  // we start with maximum = 0
	for _, cities := range premutations {

		// cities [London Belfast Dublin]
		dist := uint(0)
		for i := 0; i < len(cities)-1; i++ {
			dist += distances[cities[i]][cities[i+1]]
		}
		if dist < min {
			min = dist
		}
		if dist > max {
			max = dist
		}
	}

	return min, max
}
