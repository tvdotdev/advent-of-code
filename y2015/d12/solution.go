package d12

import (
	"encoding/json"
	"fmt"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 12, Solve)
}

func Solve(input string) {

	var jsonData interface{} // in go, interface{} = any
	json.Unmarshal([]byte(input), &jsonData)

	// fmt.Printf("%v", jsonData)

	part1 := sumAllNumbers(jsonData, false)
	fmt.Printf("Part 1: sum all %d\n", part1)

	part2 := sumAllNumbers(jsonData, true)
	fmt.Printf("Part 2: sum all, ignoring red values %d", part2)
}

func sumAllNumbers(jsonData interface{}, checkForRed bool) int {
	sum := 0

	// array
	// json file
	// mix
	// number
	// string -> parsed automatically as valu for value
	switch value := jsonData.(type) {

	// array
	case []interface{}:
		for _, v := range value {
			sum += sumAllNumbers(v, checkForRed)
		}

		// json object -> check for red value
	case map[string]interface{}:

		if checkForRed {
			for _, v := range value {
				// we have a recursion termination condition
				if v == "red" {
					return 0
				}
			}
		}
		for _, v := range value {
			sum += sumAllNumbers(v, checkForRed)
		}

		// termination of the recursion
	case int:
		sum += value

		// termination of the recursion
	case float64:
		sum += int(value)
	}

	return sum
}
