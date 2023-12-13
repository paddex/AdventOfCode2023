package p1

import (
	"strings"

	"paddex.net/aoc/types"
)

func P1(app types.App) int {
	patterns := strings.Split(app.Input, "\n\n")

	sum := 0
	for _, p := range patterns {
		sum += processPattern(p)
	}

	return sum
}

func processPattern(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	isSymmetric, numLinesAbove := checkSymmetry(lines)

	if isSymmetric {
		return numLinesAbove * 100
	} else {
		inp := turnInput(lines)
		isSymmetric, numLinesAbove := checkSymmetry(inp)
		if isSymmetric {
			return numLinesAbove
		}
	}
	return 0
}

func checkSymmetry(input []string) (bool, int) {
	idx := -1
	symmetric := false
	for i, l := range input {
		if i >= len(input)-1 {
			break
		}
		if l == input[i+1] {
			symmetric = true
			j := 1
			for i-j >= 0 && i+j+1 < len(input) {
				if input[i-j] != input[i+j+1] {
					symmetric = false
					break
				}
				j++
			}
			if symmetric {
				idx = i + 1
				break
			}
		}
	}

	return symmetric, idx
}

func turnInput(input []string) []string {
	newInp := make([]string, 0)

	for i := 0; i < len(input[0]); i++ {
		str := ""
		for j := 0; j < len(input); j++ {
			str += string(input[j][i])
		}
		newInp = append(newInp, str)
	}

	return newInp
}
