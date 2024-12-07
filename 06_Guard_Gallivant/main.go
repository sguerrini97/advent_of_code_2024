package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Vec2 struct {
	I int
	J int
}

type Guard struct {
	Position  Vec2
	Direction Vec2
}

func turn_right(p Vec2) Vec2 {
	turned := Vec2{0, 0}

	if p.I == 1 {
		turned.J = -1
	} else if p.J == -1 {
		turned.I = -1
	} else if p.I == -1 {
		turned.J = 1
	} else if p.J == 1 {
		turned.I = 1
	}

	return turned
}

func parse(input_path *string) ([][]rune, *Guard) {
	// open input file and check for errors
	input_file, err := os.Open(*input_path)
	if err != nil {
		log.Fatalf("Failed to open input file \"%s\"\nError: %s\n", *input_path, err)
	}

	defer input_file.Close()

	// prepare map of the area and guard info
	area := make([][]rune, 0)
	var guard *Guard = nil

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	i := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		area = append(area, line)

		if guard == nil {
			for j, r := range line {
				switch r {
				case 'v':
					guard = &Guard{Position: Vec2{i, j}, Direction: Vec2{1, 0}}
				case '<':
					guard = &Guard{Position: Vec2{i, j}, Direction: Vec2{0, -1}}
				case '^':
					guard = &Guard{Position: Vec2{i, j}, Direction: Vec2{-1, 0}}
				case '>':
					guard = &Guard{Position: Vec2{i, j}, Direction: Vec2{0, 1}}
				}
			}
			i += 1
		}
	}

	return area, guard
}

func out_of_area(area [][]rune, position Vec2) bool {
	if position.I < 0 || position.I >= len(area) {
		return true
	}

	if position.J < 0 || position.J >= len(area[position.I]) {
		return true
	}

	return false
}

func is_obstacle(area [][]rune, position Vec2) bool {
	r := area[position.I][position.J]

	switch r {
	case '#', 'O':
		return true
	default:
		return false
	}
}

func step(position Vec2, direction Vec2) Vec2 {
	return Vec2{I: position.I + direction.I, J: position.J + direction.J}
}

func part1(area [][]rune, guard *Guard) int {
	steps := 0

	for true {
		// increment location counter if this is a new location
		if area[guard.Position.I][guard.Position.J] != 'X' {
			steps += 1

			// mark current location as visited
			area[guard.Position.I][guard.Position.J] = 'X'
		}

		// compute next position in the area
		next_position := step(guard.Position, guard.Direction)
		if out_of_area(area, next_position) {
			break
		}

		if is_obstacle(area, next_position) {
			guard.Direction = turn_right(guard.Direction)
			continue
		}

		// move the guard
		guard.Position = next_position
	}

	return steps
}

func obstacle_on_path(area [][]rune, start Vec2, direction Vec2) *Vec2 {
	position := start
	for true {
		position = step(position, direction)

		if out_of_area(area, position) {
			break
		}

		// check if we stepped into an obstacle
		if is_obstacle(area, position) {
			return &position
		}
	}

	return nil
}

func same_position(a Vec2, b Vec2) bool {
	return a.I == b.I && a.J == b.J
}

func check_loop(area [][]rune, obstacle Vec2, start Vec2, direction Vec2) bool {
	type key struct {
		PI int
		PJ int
		DI int
		DJ int
	}

	// map to track movement
	visited := make(map[key]bool)

	position := start
	for true {
		k := key{position.I, position.J, direction.I, direction.J}

		// check if we already visited this position in this direction
		v, back_on_the_same_path := visited[k]
		if back_on_the_same_path && v {
			return true
		}

		// mark current position as visited in the current direction
		visited[k] = true

		next_position := step(position, direction)
		if out_of_area(area, next_position) {
			break
		}

		if is_obstacle(area, next_position) || same_position(next_position, obstacle) {
			direction = turn_right(direction)
			continue
		}

		position = next_position
	}

	return false
}

func part2(area [][]rune, guard *Guard) int {
	found := 0

	position := guard.Position
	direction := guard.Direction

	blacklist := make(map[Vec2]bool)

	for true {
		// we can't place an obstacle were we already
		// this also takes care of the guard initial position
		blacklist[position] = true

		right := turn_right(direction)
		next_position := step(position, direction)

		if out_of_area(area, next_position) {
			break
		}

		if is_obstacle(area, next_position) {
			direction = right
			continue
		}

		_, blacklisted := blacklist[next_position]
		if !blacklisted {
			// if we have an obstacle on the right
			// it makes sense to try to add an obstacle in front of us
			obstacle := obstacle_on_path(area, position, right)
			if obstacle != nil {
				if check_loop(area, next_position, position, right) {
					found += 1
					// exclude this position from future obstacles
					blacklist[next_position] = true
				}
			}
		}

		position = next_position
	}

	return found
}

func main() {
	// get input file from args
	part := flag.Int("part", 0, "which challenge part to solve (1 or 2)")
	input_path := flag.String("input", "input.txt", "input file path")
	flag.Parse()

	if *part < 1 || *part > 2 {
		log.Fatalf("Invalid challenge part %d", *part)
	}

	area, guard := parse(input_path)

	result := -1
	if *part == 1 {
		result = part1(area, guard)
	} else if *part == 2 {
		result = part2(area, guard)
	}

	fmt.Printf("Solution: %d\n", result)
}
