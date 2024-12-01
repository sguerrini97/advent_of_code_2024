package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// get input file from args
	input_path := flag.String("input", "input.txt", "input file path")
	flag.Parse()

	// open input file and check for errors
	input_file, err := os.Open(*input_path)
	if err != nil {
		log.Fatalf("Failed to open input file \"%s\"\nError: %s\n", *input_path, err)
	}

	defer input_file.Close()

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

	fmt.Println(sum)
}
