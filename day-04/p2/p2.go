package p2

import (
	"slices"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

func P2(app types.App) int {
	lines := strings.Split(app.Input, "\n")
	numLines := len(lines) - 1

	cardCount := make(map[int]int)
	values := make(map[int]int)

	for i := 0; i < numLines; i++ {
		cardCount[i] = 1
		values[i] = -1
	}

	for i, line := range lines {
		if line == "" {
			continue
		}
		val := processLine(line)
		for x := 0; x < val; x++ {
			cardCount[i+(x+1)] += cardCount[i]
		}
	}

	count := 0
	for i := 0; i < numLines; i++ {
		count += cardCount[i]
	}
	return count
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

	val := 0
	for _, cardNo := range cardInts {
		if slices.Contains(winningInts, cardNo) {
			val += 1
		}
	}

	return val
}
