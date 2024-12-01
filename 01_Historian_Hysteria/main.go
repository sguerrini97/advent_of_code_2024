package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func part1(input_file *os.File) int {
	ids1 := make([]int, 0)
	ids2 := make([]int, 0)

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		var id1 int
		var id2 int

		fmt.Sscanf(scanner.Text(), "%d %d", &id1, &id2)

		ids1 = append(ids1, id1)
		ids2 = append(ids2, id2)
	}

	// sort the slices
	sort.Sort(sort.IntSlice(ids1))
	sort.Sort(sort.IntSlice(ids2))

	// compute sum of sitances
	sum := 0
	for i := range ids1 {
		distance := ids2[i] - ids1[i]
		if distance < 0 {
			sum += -distance
		} else {
			sum += distance
		}
	}

	return sum
}

func part2(input_file *os.File) int {
	good_keys := make(map[int]bool)
	weigths := make(map[int]int)

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		var id1 int
		var id2 int

		fmt.Sscanf(scanner.Text(), "%d %d", &id1, &id2)

		// mark keys that appear in list 1
		good_keys[id1] = true

		// count how many times keys appear in list 2
		current := weigths[id2]
		weigths[id2] = current + 1
	}

	// compute similarity score
	similarity := 0
	for key, good := range good_keys {
		if !good {
			continue
		}

		similarity += key * weigths[key]
	}

	return similarity
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

	fmt.Println(result)
}
