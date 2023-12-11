package p1

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

func P1(app types.App) int {
	universe := getUniverse(app.Input)
	universe = expandUniverse(universe)
	universe, galaxyCount, galaxyMap := numberGalaxies(universe)

	allPairs := make([]pair, 0)
	for f := 1; f <= galaxyCount; f++ {
		for t := f + 1; t <= galaxyCount; t++ {
			allPairs = append(allPairs, pair{f, t})
		}
	}

	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {
			fmt.Printf("%s", universe[y][x])
		}
		fmt.Printf("\n")
	}

	sum := 0
	fmt.Printf("PAIRS: %v", allPairs)
	for _, pair := range allPairs {
		sum += getDistance(pair, galaxyMap)
	}

	return sum
}

func getDistance(pair pair, galaxyMap map[int]loc) int {
	steps := 0

	x1 := galaxyMap[pair.from].x
	x2 := galaxyMap[pair.to].x
	if x2 < x1 {
		xtmp := x2
		x2 = x1
		x1 = xtmp
	}
	y1 := galaxyMap[pair.from].y
	y2 := galaxyMap[pair.to].y
	if y2 < y1 {
		ytmp := y2
		y2 = y1
		y1 = ytmp
	}

	steps += x2 - x1
	steps += y2 - y1

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

	for c, idx := range rowsToExpand {
		r := make([]string, len(expRow))
		copy(r, expRow)

		universe = insertAt[[]string](universe, idx+c, r)
	}

	for y := 0; y < len(universe); y++ {
		for c, idx := range colsToExpand {
			universe[y] = insertAt(universe[y], idx+c, ".")
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
