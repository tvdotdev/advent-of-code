package d14

import (
	"fmt"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 14, Solve)
}

type Reindeer struct {
	name     string
	flySpeed int
	flyFor   int
	restFor  int
	timeline []int
	points   int
}

func Solve(input string) {

	reindeers := make([]Reindeer, 0)
	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		reindeers = append(reindeers, parse(line))

	}

	part1 := 0

	seconds := 1000
	for i := range len(reindeers) {
		reindeers[i].timeline = simulateTimeline(reindeers[i], seconds)
		if part1 < reindeers[i].timeline[seconds] {
			part1 = reindeers[i].timeline[seconds]
		}
		// fmt.Printf("%s : %d\n", reindeers[i].name, reindeers[i].timeline[1000])
	}

	fmt.Printf("Part 1: won with %d\n", part1)

}

func parse(line string) Reindeer {
	name, flySpeed, flyFor, restFor := "", 0, 0, 0

	if num, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &flySpeed, &flyFor, &restFor); num != 4 || err != nil {
		panic(err)
	}

	return Reindeer{
		name:     name,
		flySpeed: flySpeed,
		flyFor:   flyFor,
		restFor:  restFor,
		timeline: []int{0},
		points:   0,
	}

}

func simulateTimeline(reindeer Reindeer, seconds int) []int {

	timeline := reindeer.timeline

	// second 0 / start already set during parsing

	// cycle = fly + rest -> fly + rest -> fly + rest

	crossedSoFar := 0
	cycle := reindeer.flyFor + reindeer.restFor
	for i := 1; i <= seconds; i++ {
		// for the first 10 entries
		if i <= reindeer.flyFor {
			crossedSoFar += reindeer.flySpeed
			timeline = append(timeline, crossedSoFar)
			// 10 -> 137 for comet, rest for 127 seconds
		} else if i < cycle {
			timeline = append(timeline, crossedSoFar)
		}

		if i >= cycle {
			pos := i % cycle
			// 0 -> 9 (10 entried for comet,)
			if pos < reindeer.flyFor {
				crossedSoFar += reindeer.flySpeed
			}
			timeline = append(timeline, crossedSoFar)
		}
	}

	return timeline
}
