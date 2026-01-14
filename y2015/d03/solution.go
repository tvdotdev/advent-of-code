package d03

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) move(dir rune) Point {
	switch dir {
	case '^':
		return Point{p.X, p.Y + 1}
	case '>':
		return Point{p.X + 1, p.Y}
	case 'v':
		return Point{p.X, p.Y - 1}
	case '<':
		return Point{p.X - 1, p.Y}
	}

	return p
}

func Solve(input string) {

	visited := visitedHouse(input)
	fmt.Printf("Part 1: visited houses %d\n", visited)

	visited2 := visitedHouse2(input)
	fmt.Printf("Part 2: visited 2 houses %d\n", visited2)
}

func visitedHouse(input string) int {
	visited := make(map[Point]bool)
	santaPosition := Point{0, 0}

	visited[santaPosition] = true

	for _, dir := range input {
		santaPosition = santaPosition.move(dir)
		visited[santaPosition] = true
	}

	return len(visited)
}

func visitedHouse2(input string) int {
	visited := make(map[Point]bool)
	santaPosition := Point{0, 0}
	roboPosition := Point{0, 0}

	visited[santaPosition] = true
	visited[roboPosition] = true

	// santa goes on event: 0 , 2, 4, 6
	// robo goes on 1, 3, 5 , 7 .. .
	for index, dir := range input {
		if index%2 == 0 {
			santaPosition = santaPosition.move(dir)
			visited[santaPosition] = true
		} else {
			roboPosition = roboPosition.move(dir)
			visited[roboPosition] = true
		}

	}

	return len(visited)
}
