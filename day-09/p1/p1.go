package p1

import (
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

func P1(app types.App) int {
	lines := strings.Split(app.Input, "\n")
	lines = lines[:len(lines)-1]

	sum := 0
	for _, line := range lines {
		arr := getIntArr(line)
		val := getNextValue(arr)
		sum += val
	}

	return sum
}

func getIntArr(input string) []int {
	numStrs := strings.Fields(input)
	nums := make([]int, 0)
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	return nums
}

func getNextValue(input []int) int {
	allNull := true
	for _, val := range input {
		if val != 0 {
			allNull = false
			break
		}
	}
	if allNull {
		return 0
	}

	nextRow := make([]int, 0)
	for i := 0; i < len(input)-1; i++ {
		nextRow = append(nextRow, input[i+1]-input[i])
	}

	val := input[len(input)-1] + getNextValue(nextRow)
	return val
}
