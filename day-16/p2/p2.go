package p2

import (
	"fmt"
	"strings"

	"paddex.net/aoc/types"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func P2(app types.App) int {
	grid := getGrid(app.Input)

	maxSum := 0
	for x := 1; x < len(grid[0])-1; x++ {
		s1 := getSum(x, 1, grid, DOWN)
		s2 := getSum(x, len(grid)-2, grid, UP)

		if s1 > maxSum {
			maxSum = s1
		}
		if s2 > maxSum {
			maxSum = s2
		}

	}
	for y := 1; y < len(grid)-1; y++ {
		s1 := getSum(1, y, grid, RIGHT)
		s2 := getSum(len(grid[0])-2, y, grid, LEFT)

		if s1 > maxSum {
			maxSum = s1
		}
		if s2 > maxSum {
			maxSum = s2
		}

	}
	return maxSum
}

func getSum(x, y int, grid [][]string, dir int) int {
	energized := make([][]bool, len(grid)-2)
	for i := range energized {
		energized[i] = make([]bool, len(grid[i])-2)
	}
	cache := make(map[string]bool)

	ray(x, y, grid, energized, cache, dir)

	sum := 0
	for y := range energized {
		for x := range energized[y] {
			// var p string
			if energized[y][x] {
				// p = "#"
				sum++
			} else {
				// p = "."
			}
			// fmt.Printf("%s", p)
		}
		// fmt.Println()
	}
	return sum
}

func ray(x, y int, grid [][]string, energized [][]bool, cache map[string]bool, dir int) {
	var xnext, ynext int

	switch grid[y][x] {
	case "\\":
		switch dir {
		case UP:
			dir = LEFT
		case DOWN:
			dir = RIGHT
		case LEFT:
			dir = UP
		case RIGHT:
			dir = DOWN
		default:
		}
	case "/":
		switch dir {
		case UP:
			dir = RIGHT
		case DOWN:
			dir = LEFT
		case LEFT:
			dir = DOWN
		case RIGHT:
			dir = UP
		default:
		}
	case "-":
		switch dir {
		case UP:
			ray(x+1, y, grid, energized, cache, RIGHT)
			ray(x-1, y, grid, energized, cache, LEFT)
			return
		case DOWN:
			ray(x+1, y, grid, energized, cache, RIGHT)
			ray(x-1, y, grid, energized, cache, LEFT)
		default:
		}
	case "|":
		switch dir {
		case RIGHT:
			ray(x, y+1, grid, energized, cache, DOWN)
			ray(x, y-1, grid, energized, cache, UP)
			return
		case LEFT:
			ray(x, y+1, grid, energized, cache, DOWN)
			ray(x, y-1, grid, energized, cache, UP)
		default:
		}
	}

	for {
		energized[y-1][x-1] = true
		cacheString := fmt.Sprintf("%d,%d-%d", x, y, dir)
		if _, ok := cache[cacheString]; ok {
			return
		}

		cache[cacheString] = true

		if dir == UP {
			xnext = x
			ynext = y - 1
		} else if dir == RIGHT {
			xnext = x + 1
			ynext = y
		} else if dir == DOWN {
			xnext = x
			ynext = y + 1
		} else if dir == LEFT {
			xnext = x - 1
			ynext = y
		}

		nSym := grid[ynext][xnext]
		x = xnext
		y = ynext
		switch nSym {
		case ".":
		case "\\":
			if dir == RIGHT {
				dir = DOWN
			} else if dir == LEFT {
				dir = UP
			} else if dir == UP {
				dir = LEFT
			} else if dir == DOWN {
				dir = RIGHT
			}
		case "/":
			if dir == RIGHT {
				dir = UP
			} else if dir == LEFT {
				dir = DOWN
			} else if dir == UP {
				dir = RIGHT
			} else if dir == DOWN {
				dir = LEFT
			}
		case "-":
			if dir == UP || dir == DOWN {
				ray(xnext, ynext, grid, energized, cache, LEFT)
				ray(xnext, ynext, grid, energized, cache, RIGHT)
				return
			}
		case "|":
			if dir == LEFT || dir == RIGHT {
				ray(xnext, ynext, grid, energized, cache, UP)
				ray(xnext, ynext, grid, energized, cache, DOWN)
				return
			}
		case "+":
			return
		default:
			return
		}
	}
}

func getGrid(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]string, 0)
	pady := make([]string, 0)
	for i := 0; i < len(lines[0])+2; i++ {
		pady = append(pady, "+")
	}
	grid = append(grid, pady)
	for _, l := range lines {
		row := make([]string, 0)
		row = append(row, "+")
		for _, r := range l {
			row = append(row, string(r))
		}
		row = append(row, "+")
		grid = append(grid, row)
	}
	grid = append(grid, pady)

	return grid
}
