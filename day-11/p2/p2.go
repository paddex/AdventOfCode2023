package p2

import (
	"fmt"
	"strings"

	"paddex.net/aoc/types"
)

type pair struct {
	from int
	to   int
}

type loc struct {
	x int
	y int
}

func P2(app types.App) int {
	universe := getUniverse(app.Input)
	universe = expandUniverse(universe)
	universe, galaxyCount, galaxyMap := numberGalaxies(universe)

	allPairs := make([]pair, 0)
	for f := 1; f <= galaxyCount; f++ {
		for t := f + 1; t <= galaxyCount; t++ {
			allPairs = append(allPairs, pair{f, t})
		}
	}

	/*
		for y := 0; y < len(universe); y++ {
			for x := 0; x < len(universe[y]); x++ {
				fmt.Printf("%s", universe[y][x])
			}
			fmt.Printf("\n")
		}
	*/

	sum := 0
	for _, pair := range allPairs {
		sum += getDistance(pair, galaxyMap, universe)
	}

	return sum
}

func getDistance(pair pair, galaxyMap map[int]loc, universe [][]string) int {
	steps := 0

	x1 := galaxyMap[pair.from].x
	x2 := galaxyMap[pair.to].x
	xStart := x1
	if x2 < x1 {
		xtmp := x2
		x2 = x1
		x1 = xtmp
		xStart = x1
	}
	y1 := galaxyMap[pair.from].y
	y2 := galaxyMap[pair.to].y
	yStart := y1
	if y2 < y1 {
		ytmp := y2
		y2 = y1
		y1 = ytmp
		yStart = y1
	}

	for x1 < x2 {
		if (universe[yStart][x1] != "*") && (universe[yStart][x1] != "$") {
			steps++
		} else {
			steps += 1000000
		}
		x1++
	}
	for y1 < y2 {
		if (universe[y1][xStart] != "+") && (universe[y1][xStart] != "$") {
			steps++
		} else {
			steps += 1000000
		}
		y1++
	}

	return steps
}

func numberGalaxies(universe [][]string) ([][]string, int, map[int]loc) {
	count := 1
	galaxyMap := make(map[int]loc)
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {
			if universe[y][x] == "#" {
				num := fmt.Sprintf("%d", count)
				universe[y][x] = num
				galaxyMap[count] = loc{x, y}
				count++
			}
		}
	}

	return universe, count - 1, galaxyMap
}

func expandUniverse(universe [][]string) [][]string {
	colsToExpand := make([]int, 0)
	rowsToExpand := make([]int, 0)

	for x := 0; x < len(universe[0]); x++ {
		expand := true
		for y := 0; y < len(universe); y++ {
			if universe[y][x] != "." {
				expand = false
				break
			}
		}
		if expand {
			colsToExpand = append(colsToExpand, x)
		}
	}

	for y := 0; y < len(universe); y++ {
		expand := true
		for x := 0; x < len(universe[y]); x++ {
			if universe[y][x] != "." {
				expand = false
				break
			}
		}
		if expand {
			rowsToExpand = append(rowsToExpand, y)
		}
	}

	expRow := make([]string, len(universe[0]))
	copy(expRow, universe[rowsToExpand[0]])

	for _, idx := range rowsToExpand {
		r := make([]string, len(expRow))
		for i := 0; i < len(r); i++ {
			r[i] = "+"
		}
		universe[idx] = r

	}

	for y := 0; y < len(universe); y++ {
		for _, idx := range colsToExpand {
			if universe[y][idx] == "." {
				universe[y][idx] = "*"
			} else if universe[y][idx] == "+" {
				universe[y][idx] = "$"
			}
		}
	}

	return universe
}

func getUniverse(input string) [][]string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	universe := make([][]string, 0)
	for _, line := range lines {
		uniLine := make([]string, 0)
		for _, r := range line {
			uniLine = append(uniLine, string(r))
		}
		universe = append(universe, uniLine)
	}
	return universe
}

func insertAt[T any](a []T, index int, value T) []T {
	n := len(a)
	if index < 0 {
		index = (index%n + n) % n
	}
	switch {
	case index == n: // nil or empty slice or after last element
		return append(a, value)

	case index < n: // index < len(a)
		a = append(a[:index+1], a[index:]...)
		a[index] = value
		return a

	case index < cap(a): // index > len(a)
		a = a[:index+1]
		var zero T
		for i := n; i < index; i++ {
			a[i] = zero
		}
		a[index] = value
		return a

	default:
		b := make([]T, index+1) // malloc
		if n > 0 {
			copy(b, a)
		}
		b[index] = value
		return b
	}
}
