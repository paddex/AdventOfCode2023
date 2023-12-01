package p2

import (
	"fmt"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

var (
	toTest = [19]string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	numMap = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func P2(app types.App) int {
	nums := make([]int, 0)
	for _, line := range strings.Split(app.Input, "\n") {
		if !(len(line) > 0) {
			continue
		}
		num, err := getnum(line, app)
		if err != nil {
			fmt.Println(err)
			panic("Something went wrong")
		}
		app.Logger.Debug("NUMSTR", "value", num)
		nums = append(nums, num)
	}
	// fmt.Println(len(nums))
	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}

func getnum(line string, app types.App) (int, error) {
	first := ""
	last := ""
	indexFirst := 999999
	indexLast := -1

	app.Logger.Debug("LINE", "value", line)
	for _, test := range toTest {
		idx := strings.Index(line, test)
		lidx := strings.LastIndex(line, test)
		if idx < 0 {
			continue
		}
		if idx < indexFirst {
			indexFirst = idx
			if len(test) > 1 {
				first = numMap[test]
			} else {
				first = test
			}
		}
		if lidx > indexLast {
			indexLast = lidx
			if len(test) > 1 {
				last = numMap[test]
			} else {
				last = test
			}
		}
	}

	numstr := first + last
	return strconv.Atoi(numstr)
}
