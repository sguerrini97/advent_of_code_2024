package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func part1(input_file *os.File) int {
	sum := 0
	mul_regexp := regexp.MustCompile("mul\\([0-9]{1,},[0-9]{1,}\\)")

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		// find all good mul instructions
		muls := mul_regexp.FindAllString(scanner.Text(), -1)
		for _, mul := range muls {
			var x int
			var y int

			// extract operands
			fmt.Sscanf(mul, "mul(%d,%d)", &x, &y)

			sum += x * y
		}
	}

	return sum
}

func part2(input_file *os.File) int {
	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

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

	// open input file and check for errors
	input_file, err := os.Open(*input_path)
	if err != nil {
		log.Fatalf("Failed to open input file \"%s\"\nError: %s\n", *input_path, err)
	}

	defer input_file.Close()

	result := -1
	if *part == 1 {
		result = part1(input_file)
	} else if *part == 2 {
		result = part2(input_file)
	}

	fmt.Printf("Solution: %d\n", result)
}
