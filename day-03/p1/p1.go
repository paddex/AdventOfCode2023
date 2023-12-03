package p1

import (
	"fmt"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type number struct {
	x   int
	y   []int
	val int
}

type point struct {
	x int
	y int
}

// var nonsymbols = [11]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "."}
var (
	nonsymbols = "1234567890."
	nums       = "0123456789"
	size       = 0
)

func P1(app types.App) int {
	matrix := make([][]string, 0)

	lines := strings.Split(app.Input, "\n")
	size = len(lines) - 1

	for _, line := range lines {
		row := make([]string, 0)
		for _, rune := range line {
			row = append(row, string(rune))
		}
		matrix = append(matrix, row)
	}

	numbers := findNumbers(matrix)

	sum := 0
	for _, num := range numbers {

		toCheck := findToCheck(num)

		valid := checkIfSymbol(toCheck, matrix)

		if valid {
			sum += num.val
		}
	}
	return sum
}

func findNumbers(matrix [][]string) []number {
	numbers := make([]number, 0)
	isNum := false
	numy := make([]int, 0)
	numval := ""
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if strings.Contains(nums, matrix[x][y]) {
				isNum = true
				numy = append(numy, y)
				numval += matrix[x][y]
			} else {
				if isNum {
					isNum = false
					numv, _ := strconv.Atoi(numval)
					num := number{x: x, y: numy, val: numv}
					numbers = append(numbers, num)
					numy = make([]int, 0)
					numval = ""
				}
			}
		}
		if isNum {
			isNum = false
			numv, _ := strconv.Atoi(numval)
			num := number{x: x, y: numy, val: numv}
			numbers = append(numbers, num)
			numy = make([]int, 0)
			numval = ""
		}
	}

	return numbers
}

func findToCheck(num number) []point {
	fmt.Printf("%v\n", num)
	toCheck := make([]point, 0)

	starty := num.y[0]
	endy := num.y[len(num.y)-1]

	if starty > 0 {
		if num.x > 0 {
			toCheck = append(toCheck, point{num.x - 1, starty - 1})
		}
		toCheck = append(toCheck, point{num.x, starty - 1})
		if num.x < size-1 {
			toCheck = append(toCheck, point{num.x + 1, starty - 1})
		}
	}
	for i := 0; i < len(num.y); i++ {
		if num.x > 0 {
			toCheck = append(toCheck, point{num.x - 1, starty + i})
		}
		if num.x < size-1 {
			toCheck = append(toCheck, point{num.x + 1, starty + i})
		}
	}
	if endy < size-1 {
		if num.x > 0 {
			toCheck = append(toCheck, point{num.x - 1, endy + 1})
		}
		toCheck = append(toCheck, point{num.x, endy + 1})
		if num.x < size-1 {
			toCheck = append(toCheck, point{num.x + 1, endy + 1})
		}
	}

	return toCheck
}

func checkIfSymbol(toCheck []point, matrix [][]string) bool {
	valid := false
	for _, p := range toCheck {
		fmt.Printf("%v", p)
		if !strings.Contains(nonsymbols, matrix[p.x][p.y]) {
			valid = true
		}
	}
	fmt.Println()

	return valid
}
