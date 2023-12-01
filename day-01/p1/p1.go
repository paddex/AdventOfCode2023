package p1

import (
	"slices"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

var numstrings = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func P1(app types.App) int {
	nums := make([]int, 0)
	for _, line := range strings.Split(app.Input, "\n") {
		if !(len(line) > 0) {
			continue
		}
		num, err := getnum(line, app)
		if err != nil {
			app.Logger.Error(err.Error())
			panic("Something went wrong")
		}
		nums = append(nums, num)

	}
	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}

func getnum(input string, app types.App) (int, error) {
	app.Logger.Debug("Line", "value", input)
	isFirst := true
	var first, last string
	for _, char := range input {
		if !(slices.Contains(numstrings, string(char))) {
			continue
		}
		if isFirst {
			first = string(char)
			last = string(char)
			isFirst = false
		} else {
			last = string(char)
		}
	}

	numstr := first + last
	app.Logger.Debug("NumString", "value", numstr)

	return strconv.Atoi(numstr)
}
