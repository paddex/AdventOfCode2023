package p1

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	"paddex.net/aoc/types"
)

func P1(app types.App) int {
	lines := strings.Split(app.Input, "\n")

	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		val := processLine(line)
		sum += val
	}

	return sum
}

func processLine(line string) int {
	data := strings.TrimSpace(strings.Split(line, ":")[1])

	winningNosStr := strings.TrimSpace(strings.Split(data, "|")[0])
	cardNumbersStr := strings.TrimSpace(strings.Split(data, "|")[1])

	winningNos := strings.Split(winningNosStr, " ")
	cardNos := strings.Split(cardNumbersStr, " ")

	winningInts := make([]int, 0)
	cardInts := make([]int, 0)

	for _, wN := range winningNos {
		if wN == "" {
			continue
		}
		wI, _ := strconv.Atoi(wN)
		winningInts = append(winningInts, wI)
	}

	for _, cN := range cardNos {
		if cN == "" {
			continue
		}
		cI, _ := strconv.Atoi(cN)
		cardInts = append(cardInts, cI)
	}

	fmt.Printf("%v\n", winningInts)
	fmt.Printf("%v\n", cardInts)

	val := 0
	first := true
	for _, cardNo := range cardInts {
		if slices.Contains(winningInts, cardNo) {
			if first {
				first = false
				val = 1
				continue
			}
			val *= 2
		}
	}

	return val
}
