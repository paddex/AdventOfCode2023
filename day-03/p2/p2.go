package p2

import (
	"fmt"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type star struct {
	x int
	y int
}

type gear struct {
	x    int
	y    int
	nums []int
}

type number struct {
	str   string
	val   int
	start int
	end   int
}

var (
	size    = 0
	numbers = "0123456789"
)

func P2(app types.App) int {
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

	stars := findStars(matrix)

	gears := make([]gear, 0)
	for _, star := range stars {
		nums := findNumbers(star, matrix)

		if len(nums) == 2 {
			gears = append(gears, gear{x: star.x, y: star.y, nums: []int{nums[0].val, nums[1].val}})
		}
	}

	sum := 0
	for _, gear := range gears {
		sum += gear.nums[0] * gear.nums[1]
	}

	return sum
}

func findStars(matrix [][]string) []star {
	stars := make([]star, 0)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if matrix[x][y] == "*" {
				stars = append(stars, star{x: x, y: y})
			}
		}
	}
	return stars
}

func findNumbers(star star, matrix [][]string) []number {
	nums := make([]number, 0)

	fmt.Printf("STAR: %v\n", star)
	for x := -1; x <= 1; x++ {
		if star.x+x < 0 || star.x+x > size {
			continue
		}
		for y := -1; y <= 1; y++ {
			if star.y+y < 0 || star.y+y > size {
				continue
			}
			if strings.Contains("0123456789", matrix[star.x+x][star.y+y]) {
				num := getNumber(star.x+x, star.y+y, matrix)
				contained := false
				for _, oldnum := range nums {
					if oldnum == num {
						contained = true
					}
				}
				if !contained {
					fmt.Printf("NUM: %v\n", num)
					nums = append(nums, num)
				}
			}
		}
	}

	return nums
}

func getNumber(x int, y int, matrix [][]string) number {
	fmt.Printf("%d %d\n\n", x, y)
	tmpy := y
	var start int
	var end int

	val := matrix[x][y]
	if tmpy > 0 {
		for strings.Contains(numbers, matrix[x][tmpy-1]) {
			val = matrix[x][tmpy-1] + val
			tmpy = tmpy - 1
			if tmpy <= 0 {
				break
			}
		}
	}

	start = tmpy

	tmpy = y
	if tmpy < size-1 {
		for strings.Contains(numbers, matrix[x][tmpy+1]) {
			val = val + matrix[x][tmpy+1]
			tmpy = tmpy + 1
			if tmpy >= size-1 {
				break
			}
		}
	}

	end = tmpy

	numval, _ := strconv.Atoi(val)

	return number{str: val, val: numval, start: start, end: end}
}
