package main

import "fmt"

func main() {
	inputs := [][]int{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0, 1, 1, 1, 1},
	}

	for _, input := range inputs {
		result := convert(input)

		fmt.Println(fmt.Sprintf("%v # => %v", input, result))
	}
}

func convert(input []int) []string {
	length := len(input)

	cuts := make([]string, 0)
	lastCut := 0

	for i, val := range input {
		if val == 1 {
			// Ignore first cut
			if i != 0 {
				cuts = append(cuts, fmt.Sprintf("b%d", length/(i-lastCut)))
			}

			lastCut = i
		}
	}

	// Capture remaining length
	return append(cuts, fmt.Sprintf("b%d", length/(length-lastCut)))
}
