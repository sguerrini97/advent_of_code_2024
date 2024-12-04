package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func compare_buffers(a *[]rune, b *[]rune) bool {
	if len(*a) != len(*b) {
		return false
	}

	match := true

	for i := range *a {
		if (*a)[i] != (*b)[i] {
			match = false
			break
		}
	}

	return match
}

func check_word(word *[]rune, i int, j int, mat *[][]rune) int {
	side := len(*mat)
	word_len := len(*word)
	buffer := make([]rune, word_len)

	found := 0

	// check matrix boundaries
	can_go_left := j-(word_len-1) >= 0
	can_go_right := j+(word_len-1) <= side-1
	can_go_down := i+(word_len-1) <= side-1
	can_go_up := i-(word_len-1) >= 0

	if can_go_down {
		// check down
		for w := range word_len {
			buffer[w] = (*mat)[i+w][j]
		}

		if compare_buffers(&buffer, word) {
			found += 1
		}
	}

	if can_go_up {
		// check up
		for w := range word_len {
			buffer[w] = (*mat)[i-w][j]
		}

		if compare_buffers(&buffer, word) {
			found += 1
		}
	}

	if can_go_left {
		// check left
		for w := range word_len {
			buffer[w] = (*mat)[i][j-w]
		}

		if compare_buffers(&buffer, word) {
			found += 1
		}

		if can_go_down {
			// check down diagonal
			for w := range word_len {
				buffer[w] = (*mat)[i+w][j-w]
			}

			if compare_buffers(&buffer, word) {
				found += 1
			}
		}

		if can_go_up {
			// check up diagonal
			for w := range word_len {
				buffer[w] = (*mat)[i-w][j-w]
			}

			if compare_buffers(&buffer, word) {
				found += 1
			}
		}
	}

	if can_go_right {
		// check right
		for w := range word_len {
			buffer[w] = (*mat)[i][j+w]
		}

		if compare_buffers(&buffer, word) {
			found += 1
		}

		if can_go_down {
			// check down diagonal
			for w := range word_len {
				buffer[w] = (*mat)[i+w][j+w]
			}

			if compare_buffers(&buffer, word) {
				found += 1
			}
		}

		if can_go_up {
			// check up diagonal
			for w := range word_len {
				buffer[w] = (*mat)[i-w][j+w]
			}

			if compare_buffers(&buffer, word) {
				found += 1
			}
		}
	}

	return found
}

func part1(mat *[][]rune) int {
	xmas_count := 0

	looking_for := []rune{'X', 'M', 'A', 'S'}

	for i := range *mat {
		for j := range (*mat)[i] {
			r := (*mat)[i][j]

			if r == looking_for[0] {
				found := check_word(&looking_for, i, j, mat)
				xmas_count += found
			}
		}
	}

	return xmas_count
}

func part2(mat *[][]rune) int {
	found := 0

	side := len(*mat)
	for i := range side {
		// skip first and last row
		if i == 0 || i == side-1 {
			continue
		}

		for j := range side {
			// skip first and last column
			if j == 0 || j == side-1 {
				continue
			}

			// check for X-MAS
			r := (*mat)[i][j]
			if r == 'A' {
				// we may be at the center of an X-MAS
				// check for the corners to confirm that
				ul := (*mat)[i-1][j-1]
				dr := (*mat)[i+1][j+1]
				ur := (*mat)[i-1][j+1]
				dl := (*mat)[i+1][j-1]

				if (ul == 'M' && dr == 'S') || (dr == 'M' && ul == 'S') {
					if (ur == 'M' && dl == 'S') || (dl == 'M' && ur == 'S') {
						found += 1
					}
				}
			}
		}
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

	// open input file and check for errors
	input_file, err := os.Open(*input_path)
	if err != nil {
		log.Fatalf("Failed to open input file \"%s\"\nError: %s\n", *input_path, err)
	}

	defer input_file.Close()

	var mat [][]rune

	// read file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		// convert lines to a rune matrix
		mat = append(mat, []rune(scanner.Text()))
	}

	result := -1
	if *part == 1 {
		result = part1(&mat)
	} else if *part == 2 {
		result = part2(&mat)
	}

	fmt.Printf("Solution: %d\n", result)
}
