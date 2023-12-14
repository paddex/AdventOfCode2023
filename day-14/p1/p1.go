package p1

import (
	"strings"

	"paddex.net/aoc/types"
)

func P1(app types.App) int {
	g := getGrid(app.Input)

	g = tiltNorth(g)

	/*
		for y := range g {
			for x := range g[y] {
				fmt.Printf("%s", g[y][x])
			}
			fmt.Println()
		}
	*/

	return getLoad(g)
}

func getLoad(grid [][]string) int {
	load := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" {
				load += 1 * (len(grid) - y)
			}
		}
	}
	return load
}

func tiltNorth(grid [][]string) [][]string {
	for y, row := range grid {
		for x, sym := range row {
			if sym == "O" {
				grid[y][x] = "."
				i := 0
				for i <= y {
					if grid[y-i][x] != "." {
						break
					}
					i++
				}
				grid[y-i+1][x] = "O"
			}
		}
	}

	return grid
}

func getGrid(input string) [][]string {
	grid := make([][]string, 0)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for y, l := range lines {
		L := make([]string, 0)
		for x := 0; x < len(lines[y]); x++ {
			L = append(L, string(l[x]))
		}
		grid = append(grid, L)
	}

	return grid
}

func someFunc(input string) int {
	return 0
}
