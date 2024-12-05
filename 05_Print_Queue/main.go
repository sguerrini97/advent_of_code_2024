package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func parse(input_path *string) (map[int][]int, [][]int) {
	input_file, err := os.Open(*input_path)
	if err != nil {
		log.Fatalf("Failed to open input file \"%s\"\nError: %s\n", *input_path, err)
	}

	defer input_file.Close()

	constraints := make(map[int][]int)

	// read page order section of input
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		line := scanner.Text()

		// input separator
		if len(line) == 0 {
			break
		}

		var x int
		var y int

		fmt.Sscanf(line, "%d|%d", &x, &y)

		// create a map of all the values the pages that cannot be after a given page
		_, present := constraints[y]
		if !present {
			constraints[y] = make([]int, 0)
		}
		constraints[y] = append(constraints[y], x)
	}

	updates := make([][]int, 0)

	// read updates section of the input
	for scanner.Scan() {
		update_strs := strings.Split(scanner.Text(), ",")

		update := make([]int, len(update_strs))

		for ui, update_str := range update_strs {
			value, _ := strconv.Atoi(update_str)
			update[ui] = value
		}

		updates = append(updates, update)
	}

	return constraints, updates
}

func part1(constraints map[int][]int, updates [][]int) int {
	sum := 0

	// check each update
	for _, update := range updates {
		good := true
		forbidden := make(map[int]bool)

		for _, page := range update {
			// if this page cannot appear, this update is not valid
			_, present := forbidden[page]
			if present {
				good = false
				break
			}

			// check to add restriction for pages after this one
			to_forbid, present := constraints[page]
			if present {
				for _, forbidden_page := range to_forbid {
					forbidden[forbidden_page] = true
				}
			}
		}

		if good {
			i := len(update) / 2
			sum += update[i]
		}
	}

	return sum
}

func part2(constraints map[int][]int, updates [][]int) int {
	sum := 0

	// check each update
	for _, update := range updates {
		forbidden := make(map[int]bool)

		for _, page := range update {
			// if this page cannot appear, this update is not valid
			_, present := forbidden[page]
			if present {
				// sort using our constraints
				sort.Slice(update, func(i int, j int) bool {
					a := update[i]
					b := update[j]

					_, present := constraints[b]
					if present && slices.Contains(constraints[b], a) {
						return true
					}

					return false
				})

				// we fixed this update so we can skip to the next one
				i := len(update) / 2
				sum += update[i]
				break
			}

			// check to add restriction for pages after this one
			to_forbid, present := constraints[page]
			if present {
				for _, forbidden_page := range to_forbid {
					forbidden[forbidden_page] = true
				}
			}
		}
	}

	return sum
}

func main() {
	// get input file from args
	part := flag.Int("part", 0, "which challenge part to solve (1 or 2)")
	input_path := flag.String("input", "input.txt", "input file path")
	flag.Parse()

	if *part < 1 || *part > 2 {
		log.Fatalf("Invalid challenge part %d", *part)
	}

	// parse input files into structured data
	constraints, updates := parse(input_path)

	result := -1
	if *part == 1 {
		result = part1(constraints, updates)
	} else if *part == 2 {
		result = part2(constraints, updates)
	}

	fmt.Printf("Solution: %d\n", result)
}
