package d07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tvdotdev/advent-of-code/register"
)

func init() {
	register.Register(2015, 7, Solve)
}

func Solve(input string) {
	instructions := make([]Instruction, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		// do something
		instructions = append(instructions, parse(line))
	}

	part1 := process(instructions, "a")

	fmt.Printf("Part 1: value of a %d\n", part1)

	// part 2 //
	for i, instruction := range instructions {
		if instruction.Wire == "b" {
			instructions[i] = Instruction{
				Left:  strconv.Itoa(int(part1)),
				Op:    "",
				Right: "",
				Wire:  "b",
			}
		}
	}

	part2 := process(instructions, "a")
	fmt.Printf("Part 2: value of a: %d", part2)

}

type Instruction struct {
	Left  string
	Op    string
	Right string
	Wire  string
}

func parse(line string) Instruction {
	parts := strings.Split(line, " -> ") // always two parts
	wire := parts[1]

	// tokens := strings.Split(line, " ") // ["a", "", "AND"]
	tokens := strings.Fields(parts[0])
	if len(tokens) == 1 {
		// direct assignment
		// 123 -> x
		// y -> x
		return Instruction{
			Left:  tokens[0],
			Op:    "",
			Right: "",
			Wire:  wire,
		}
	} else if len(tokens) == 2 {
		// NOT a -> y
		// NOT 1 -> c
		return Instruction{
			Left:  "",
			Op:    tokens[0],
			Right: tokens[1],
			Wire:  wire,
		}
	} else {
		// a AND y -> z
		// a LSHIFT 2 > x
		return Instruction{
			Left:  tokens[0],
			Op:    tokens[1],
			Right: tokens[2],
			Wire:  wire,
		}
	}
}

func process(instructions []Instruction, wire string) uint16 {
	// the values of processed wires
	// at first has no values
	// wires populated step by step
	wires := make(map[string]uint16)

	// until all wires has value
	for len(wires) < len(instructions) {
		for _, i := range instructions {
			if _, found := wires[i.Wire]; found {
				continue
			}

			// left and right values // processes
			leftValue, leftOk, rightValue, rightOk := 0, false, 0, false
			if i.Left == "" {
				leftOk = true
			} else if num, err := strconv.Atoi(i.Left); err == nil {
				leftValue = num
				leftOk = true
			} else {
				val, found := wires[i.Left]
				if found {
					leftValue = int(val)
				}
				leftOk = found
			}

			// right values
			if i.Right == "" {
				rightOk = true
			} else if num, err := strconv.Atoi(i.Right); err == nil {
				rightValue = num
				rightOk = true
			} else {
				val, found := wires[i.Right]
				if found {
					rightValue = int(val)
				}
				rightOk = found
			}

			// process the actual operator

			switch i.Op {

			case "":
				// 123 ->x
				// y -> s
				if leftOk {
					wires[i.Wire] = uint16(leftValue)
				}

			case "NOT":
				if rightOk {
					wires[i.Wire] = ^uint16(rightValue) // ^ is the bitwise NOT in GO
				}

			case "AND":
				// something AND someotherthing -> someotherotherthing
				if leftOk && rightOk { // logical
					wires[i.Wire] = uint16(leftValue) & uint16(rightValue) // bitwise
				}

			case "OR":
				if leftOk && rightOk {
					wires[i.Wire] = uint16(leftValue) | uint16(rightValue)
				}

			case "LSHIFT":
				if leftOk && rightOk {
					wires[i.Wire] = uint16(leftValue) << uint16(rightValue)
				}

			case "RSHIFT":
				if leftOk && rightOk {
					wires[i.Wire] = uint16(leftValue) >> uint16(rightValue)
				}
			}
		}
	}

	// for k, v := range wires {
	// 	fmt.Printf("%s : %d\n", k, v)
	// }

	if v, found := wires[wire]; found {
		return v
	}
	panic("Unable to process values of " + wire)
}
