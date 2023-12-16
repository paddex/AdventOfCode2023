package p1

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

func P1(app types.App) int {
	grid := getGrid(app.Input)
	energized := make([][]bool, len(grid)-2)
	for i := range energized {
		energized[i] = make([]bool, len(grid[i])-2)
	}
	cache := make(map[string]bool)

	var dir int
	switch grid[1][1] {
	case ".":
		dir = RIGHT
	case "\\":
		dir = DOWN
	case "/":
		dir = UP
	case "-":
		dir = RIGHT
	case "|":
		dir = DOWN
	default:
		dir = RIGHT
	}
	ray(1, 1, grid, energized, cache, dir)

	sum := 0
	for y := range energized {
		for x := range energized[y] {
			var p string
			if energized[y][x] {
				p = "#"
				sum++
			} else {
				p = "."
			}
			fmt.Printf("%s", p)
		}
		fmt.Println()
	}
	return sum
}

func ray(x, y int, grid [][]string, energized [][]bool, cache map[string]bool, dir int) {
	var xnext, ynext int

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
