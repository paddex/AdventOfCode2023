package p2

import (
	"fmt"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

var cache = make(map[state]int)

type state struct {
	springs string
	amount  string
}

func P2(app types.App) int {
	lines := strings.Split(app.Input, "\n")
	lines = lines[:len(lines)-1]

	sum := 0
	for _, line := range lines {
		possibilities := processLine(line)
		sum += possibilities
	}
	return sum
}

func setCache(springs string, amount []int, res int) int {
	am := ""
	for _, a := range amount {
		am += fmt.Sprintf("%d", a)
	}
	cache[state{springs, am}] = res
	return res
}

func processLine(input string) int {
	springs := strings.Split(input, " ")[0]
	amountsStr := strings.Split(strings.Split(input, " ")[1], ",")
	amounts := make([]int, 0)
	for _, aS := range amountsStr {
		a, _ := strconv.Atoi(aS)
		amounts = append(amounts, a)
	}

	for i := 0; i < 5; i++ {
	}

	sols := dp(springs+".", amounts)

	return sols
}

func dp(springs string, amounts []int, grouplen ...int) int {
	var grlen int
	if len(grouplen) > 0 {
		grlen = grouplen[0]
	} else {
		grlen = 0
	}

	if len(springs) == 0 {
		if len(amounts) == 0 && grlen == 0 {
			return 1
		} else {
			return 0
		}
	}

	aStr := ""
	for _, a := range amounts {
		aStr += fmt.Sprintf("%d", a)
	}
	if res, ok := cache[state{springs, aStr}]; ok {
		return res
	}

	num_solutions := 0
	symbols := make([]string, 0)
	if string(springs[0]) == "?" {
		symbols = append(symbols, ".")
		symbols = append(symbols, "#")
	} else {
		symbols = append(symbols, string(springs[0]))
	}

	for _, sym := range symbols {
		if sym == "." {
			if grlen > 0 {
				if len(amounts) > 0 && amounts[0] == grlen {
					num_solutions += dp(springs[1:], amounts[1:])
				}
			} else {
				num_solutions += dp(springs[1:], amounts)
			}
		}
		if sym == "#" {
			num_solutions += dp(springs[1:], amounts, grlen+1)
		}
	}

	return num_solutions
}
