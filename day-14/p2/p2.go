package p2

import (
	"fmt"
	"strings"

	"paddex.net/aoc/types"
)

func P2(app types.App) int {
	g1 := getGrid(app.Input)

	maxC := 10
	c := 0
	sList := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		g1 = rotate(g1)
		if c < maxC {
			fmt.Println(getString(g1))
			fmt.Println()
			c++
		}

		if i == 12 {
			fmt.Println("-----------------")
			fmt.Println("")
			fmt.Println(getString(g1))
			fmt.Println("")
			fmt.Println("-----------------")
		}
		sList[i] = getString(g1)
	}

	l, m := findCycle(sList)

	fmt.Printf("lam: %d, mu: %d\n", l, m)

	g := getGrid(app.Input)
	fmt.Println("\n\nORIGINAL G:")
	fmt.Println(getString(g))
	rc := 0
	for i := 0; i < m; i++ {
		g = rotate(g)
		rc++
	}

	for i := 0; i < l; i++ {
		g = rotate(g)
		rc++
	}

	var r int
	r = (1000000000 - m) % l
	for i := 0; i < r; i++ {
		g = rotate(g)
		rc++
	}

	fmt.Printf("ROTATION COUNT: %d\n", rc)
	fmt.Println(getString(g))

	/* for i := 0; i < 6; i++ {
		g = rotate(g)
		s := getLoad(g)
		fmt.Printf("Load after %d: %d\n", i, s)
	} */

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

func findCycle(sList []string) (int, int) {
	tc := 1
	hc := 2
	tortoise := sList[tc]
	hare := sList[hc]

	for tortoise != hare {
		tc += 1
		hc += 2
		tortoise = sList[tc]
		hare = sList[hc]
	}

	mu := 0
	tc = 0
	tortoise = sList[tc]
	for tortoise != hare {
		tc += 1
		tortoise = sList[tc]
		hc += 1
		hare = sList[hc]
		mu += 1
	}

	lam := 1
	hc = tc + 1
	hare = sList[hc]
	for tortoise != hare {
		hc += 1
		hare = sList[hc]
		lam += 1
	}

	return lam, mu
}

func getLoad(grid [][]string) int {
	load := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" {
				load += (len(grid) - y)
			}
		}
	}
	return load
}

func rotate(grid [][]string) [][]string {
	grid = tiltNorth(grid)
	grid = tiltWest(grid)
	grid = tiltSouth(grid)
	grid = tiltEast(grid)

	return grid
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

func tiltSouth(grid [][]string) [][]string {
	for y := len(grid) - 1; y > -1; y-- {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "O" {
				grid[y][x] = "."
				i := 0
				for i <= len(grid)-1-y {
					if grid[y+i][x] != "." {
						break
					}
					i++
				}
				grid[y+i-1][x] = "O"
			}
		}
	}

	return grid
}

func tiltWest(grid [][]string) [][]string {
	// for y, row := range grid {
	// 	for x, sym := range row {
	// 		if sym == "O" {
	// 			grid[y][x] = "."
	// 			i := 0
	// 			for i < len(grid[y])-x {
	// 				if grid[y][x+i] != "." {
	// 					break
	// 				}
	// 				i++
	// 			}
	// 			grid[y][x+i-1] = "O"
	// 		}
	// 	}
	// }
	for y, row := range grid {
		for x, sym := range row {
			if sym == "O" {
				grid[y][x] = "."
				i := 0
				for i <= x {
					if grid[y][x-i] != "." {
						break
					}
					i++
				}
				grid[y][x-i+1] = "O"
			}
		}
	}

	return grid
}

func tiltEast(grid [][]string) [][]string {
	for y := 0; y < len(grid); y++ {
		for x := len(grid[y]) - 1; x > -1; x-- {
			if grid[y][x] == "O" {
				grid[y][x] = "."
				i := 0
				for i < len(grid[y])-x {
					if grid[y][x+i] != "." {
						break
					}
					i++
				}
				grid[y][x+i-1] = "O"
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

func getString(input [][]string) string {
	list := make([]string, 0)
	for _, row := range input {
		list = append(list, strings.Join(row, ""))
	}

	return strings.Join(list, "\n")
}
