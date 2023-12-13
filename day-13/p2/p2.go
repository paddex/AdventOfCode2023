package p2

import (
	"strings"

	"paddex.net/aoc/types"
)

func P2(app types.App) int {
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
	diffs := 0
	for i := range input {
		if i >= len(input)-1 {
			break
		}
		diffs = 0
		for c := 0; c < len(input[i]); c++ {
			if input[i][c] != input[i+1][c] {
				diffs++
			}
		}
		if diffs < 2 {
			symmetric = true
			j := 1
			for i-j >= 0 && i+j+1 < len(input) {
				for k := 0; k < len(input[i]); k++ {
					if input[i-j][k] != input[i+j+1][k] {
						diffs++
					}
				}
				if diffs > 1 {
					symmetric = false
					break
				}
				j++
			}
			if diffs < 1 {
				symmetric = false
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
