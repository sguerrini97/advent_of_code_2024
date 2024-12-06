package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type MatrixPoint struct {
	I int
	J int
}

type Guard struct {
	Position  MatrixPoint
	Direction MatrixPoint
}

func turn_right(p *MatrixPoint) {
	if p.I == 1 {
		p.I = 0
		p.J = -1
	} else if p.J == -1 {
		p.I = -1
		p.J = 0
	} else if p.I == -1 {
		p.I = 0
		p.J = 1
	} else if p.J == 1 {
		p.I = 1
		p.J = 0
	}
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
					guard = &Guard{Position: MatrixPoint{i, j}, Direction: MatrixPoint{1, 0}}
				case '<':
					guard = &Guard{Position: MatrixPoint{i, j}, Direction: MatrixPoint{0, -1}}
				case '^':
					guard = &Guard{Position: MatrixPoint{i, j}, Direction: MatrixPoint{-1, 0}}
				case '>':
					guard = &Guard{Position: MatrixPoint{i, j}, Direction: MatrixPoint{0, 1}}
				}
			}
			i += 1
		}
	}

	return area, guard
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
		next_position := MatrixPoint{
			guard.Position.I + guard.Direction.I,
			guard.Position.J + guard.Direction.J,
		}

		// check if the guard will be out of the area
		if next_position.I < 0 || next_position.I >= len(area) ||
			next_position.J < 0 || next_position.J >= len(area[guard.Position.I]) {
			break
		}

		// check for obstacle
		r := area[next_position.I][next_position.J]
		if r == '#' {
			turn_right(&guard.Direction)
			continue
		}

		// move the guard
		guard.Position = next_position
	}

	return steps
}

func part2(area [][]rune, guard *Guard) int {

	return 0
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
