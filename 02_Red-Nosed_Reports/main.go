package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input_file *os.File) int {
	safe := 0

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		readings := strings.Split(scanner.Text(), " ")

		good := true
		increasing := false
		decreasing := false

		previous_val := 0
		for i := range readings {
			val, _ := strconv.Atoi(readings[i])

			// don't check anything on the first value
			if i == 0 {
				previous_val = val
				continue
			}

			// not safe because value changed too much
			distance := val - previous_val
			if distance > 3 || distance < -3 {
				good = false
				break
			}

			// check if we are still increasing/decreasing
			if distance > 0 {
				increasing = true
				if decreasing {
					// fail because we changed direction
					good = false
					break
				}
			} else if distance < 0 {
				decreasing = true
				if increasing {
					// fail because we changed direction
					good = false
					break
				}
			} else {
				// fail because consecutive values must change
				good = false
				break
			}

			// store current value
			previous_val = val
		}

		// if we didn't break before this test is safe
		if good {
			safe += 1
		}
	}

	return safe
}

func part2(input_file *os.File) int {
	safe := 0

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		readings := strings.Split(scanner.Text(), " ")

		good := true
		skip := -1

		// skip each step as long as we get an error
		for skip < len(readings) {
			previous_val := 0
			good = true
			increasing := false
			decreasing := false

			for i := range readings {
				if i == skip {
					// skip this step
					continue
				}

				val, _ := strconv.Atoi(readings[i])

				// don't check anything on the first value
				if i == 0 || (i == 1 && skip == 0) {
					previous_val = val
					continue
				}

				// not safe because value changed too much
				distance := val - previous_val
				if distance > 3 || distance < -3 {
					good = false
					break
				}

				// check if we are still increasing/decreasing
				if distance > 0 {
					increasing = true
					if decreasing {
						// fail because we changed direction
						good = false
						break
					}
				} else if distance < 0 {
					decreasing = true
					if increasing {
						// fail because we changed direction
						good = false
						break
					}
				} else {
					// fail because consecutive values must change
					good = false
					break
				}

				// store current value
				previous_val = val
			}

			// if we didn't break there is no problem with this test
			// and there's no need to continue
			if good {
				break
			}

			// skip the next value and check the test again
			skip += 1
		}

		// if we didn't break before this test is safe
		if good {
			safe += 1
		}
	}

	return safe
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
