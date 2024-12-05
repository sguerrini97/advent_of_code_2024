package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func parse(input_path *string) {

	input_file, err := os.Open(*input_path)
	if err != nil {
		log.Fatalf("Failed to open input file \"%s\"\nError: %s\n", *input_path, err)
	}

	defer input_file.Close()

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
	}

	// read updates section of the input
	for scanner.Scan() {
	}
}

func part1() int {

	return 0
}

func part2() int {

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

	parse(input_path)

	result := -1
	if *part == 1 {
		result = part1()
	} else if *part == 2 {
		result = part2()
	}

	fmt.Printf("Solution: %d\n", result)
}
