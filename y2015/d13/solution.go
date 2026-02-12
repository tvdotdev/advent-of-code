package d13

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 13, Solve)
}

func Solve(input string) {
	// alice -> bob -> 45
	// alice -> carol -> -50
	units := make(map[string]map[string]int)

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		parse(strings.TrimSuffix(line, "."), &units)
	}

	arrangments := premutations(slices.Collect(maps.Keys(units)))

	part1 := totalChangeInHappenis(arrangments, units)
	fmt.Printf("Part 1: total change in hapiness %d\n", part1)

	// part 2

	addMeToUnits(&units)

	arrangmentsWithMe := premutations(slices.Collect(maps.Keys(units)))

	part2 := totalChangeInHappenis(arrangmentsWithMe, units)
	fmt.Printf("Part 2: total change in hapiness with me %d\n", part2)
}

func addMeToUnits(unitsRef *map[string]map[string]int) {
	// adding my self
	(*unitsRef)["me"] = make(map[string]int)

	// add my point as 0 when I sit next to each person

	for person := range *unitsRef {
		if person == "me" {
			continue
		}
		(*unitsRef)["me"][person] = 0
		(*unitsRef)[person]["me"] = 0
	}

}
func totalChangeInHappenis(arrangments [][]string, units map[string]map[string]int) int {

	sum := 0

	for _, arrangment := range arrangments { //

		partialSum := 0

		// between the first ->  last one
		for i := 0; i < len(arrangment)-1; i++ {
			partialSum += units[arrangment[i]][arrangment[i+1]]
			partialSum += units[arrangment[i+1]][arrangment[i]]
		}

		// first + last only
		partialSum += units[arrangment[0]][arrangment[len(arrangment)-1]]
		partialSum += units[arrangment[len(arrangment)-1]][arrangment[0]]

		if partialSum > sum {
			sum = partialSum
		}
	}

	return sum
}

func parse(line string, unitsRef *map[string]map[string]int) {
	// Alice would gain 54 happiness units by sitting next to Bob.

	name1, gainOrLose, units, name2 := "", "", 0, ""

	// regexp
	// sscanf, somehow have weird issue with (.) at the end, so we need to remove it
	if num, err := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &name1, &gainOrLose, &units, &name2); num != 4 || err != nil {
		panic(err)
	}

	if (*unitsRef)[name1] == nil {
		(*unitsRef)[name1] = make(map[string]int)
	}

	sign := 1

	if gainOrLose == "lose" {
		sign = -1
	}

	// lost -> sign -1
	// gain -> sign +1

	(*unitsRef)[name1][name2] = sign * units
}

func premutations(people []string) [][]string {
	if len(people) == 1 {
		// ['Alice'] -> [ ['alice'] ]
		return [][]string{people}
	}

	results := make([][]string, 0)
	for i, person := range people {
		rem := make([]string, len(people)-1)
		// copy (to, from)
		// [:i] from the begining-> i(not include)
		// [0, 1,2,3,4,5,6,7]
		// i := 3
		// [:3] = [0, 1,2]
		// [3:] = [3...7]
		copy(rem[:i], people[:i])
		copy(rem[i:], people[i+1:])

		subPremutations := premutations(rem)

		for _, v := range subPremutations {
			// add the fixed item (person) to each sub arrangment / premutation
			// add them to the results
			results = append(results, append([]string{person}, v...))
		}
	}
	return results
}
