package p2

import (
	"fmt"
	"strings"

	"paddex.net/aoc/types"
)

const (
	WEST = iota
	NORTH
	EAST
	SOUTH
)

type (
	board [][]string
)

type loc struct {
	x int
	y int
}

type pipe struct {
	symbol string
	loc
	lastDir int
}

var loop [][]bool

func P2(app types.App) int {
	loop = make([][]bool, 0)
	b := getBoard(app.Input)
	s := findStart(b)
	ps := findStartPipes(b, s)

	steps := 0
	p := findNextPipe(b, ps[0])
	steps++
	for p.symbol != "S" {
		p = findNextPipe(b, p)
		steps++
	}

	b[s.y][s.x] = findStartingSymbol(ps)

	b2 := make([][]string, len(b))
	for i := range b {
		b2[i] = make([]string, len(b[i]))
		copy(b2[i], b[i])
	}

	enclosed := 0
	inLoop := false
	trackingF := false
	trackingL := false
	for y := 1; y < len(b)-1; y++ {
		for x := 1; x < len(b[y])-1; x++ {
			if loop[y][x] == true {
				if b[y][x] == "|" {
					inLoop = !inLoop
				}
				if b[y][x] == "F" {
					trackingF = true
				}
				if b[y][x] == "L" {
					trackingL = true
				}
				if b[y][x] == "7" {
					if trackingF {
						trackingF = false
					}
					if trackingL {
						inLoop = !inLoop
						trackingL = false
					}
				}
				if b[y][x] == "J" {
					if trackingL {
						trackingL = false
					}
					if trackingF {
						inLoop = !inLoop
						trackingF = false
					}
				}
			}
			if !loop[y][x] {
				if inLoop {
					enclosed++
					b2[y][x] = "X"
				} else {
					b2[y][x] = "O"
				}
			}
		}
		trackingF = false
		trackingL = false
		inLoop = false
	}

	for y := 1; y < len(b2)-1; y++ {
		for x := 1; x < len(b2[y])-1; x++ {
			/* var v int
			if loop[y][x] {
				v = 1
			} else {
				v = 0
			}
			fmt.Printf("%d ", v) */
			fmt.Printf("%s", b2[y][x])
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	return enclosed
}

func findStartingSymbol(ps []pipe) string {
	m := make(map[int]int)
	m[WEST] = 0
	m[NORTH] = 0
	m[EAST] = 0
	m[SOUTH] = 0
	for _, p := range ps {
		m[p.lastDir] += 1
	}

	if m[WEST] == 1 && m[EAST] == 1 {
		return "-"
	}
	if m[WEST] == 1 && m[SOUTH] == 1 {
		return "L"
	}
	if m[WEST] == 1 && m[NORTH] == 1 {
		return "F"
	}
	if m[EAST] == 1 && m[NORTH] == 1 {
		return "7"
	}
	if m[EAST] == 1 && m[SOUTH] == 1 {
		return "J"
	}
	if m[SOUTH] == 1 && m[NORTH] == 1 {
		return "|"
	}

	return ""
}

func findNextPipe(b board, p pipe) pipe {
	var l loc
	var symbol string
	var lastDir int
	if p.symbol == "-" {
		if p.lastDir == WEST {
			l = loc{p.loc.x + 1, p.loc.y}
			symbol = b[l.y][l.x]
			lastDir = WEST
		}
		if p.lastDir == EAST {
			l = loc{p.loc.x - 1, p.loc.y}
			symbol = b[l.y][l.x]
			lastDir = EAST
		}
	} else if p.symbol == "|" {
		if p.lastDir == SOUTH {
			l = loc{p.loc.x, p.loc.y - 1}
			symbol = b[l.y][l.x]
			lastDir = SOUTH
		}
		if p.lastDir == NORTH {
			l = loc{p.loc.x, p.loc.y + 1}
			symbol = b[l.y][l.x]
			lastDir = NORTH
		}
	} else if p.symbol == "7" {
		if p.lastDir == WEST {
			l = loc{p.loc.x, p.loc.y + 1}
			symbol = b[l.y][l.x]
			lastDir = NORTH
		}
		if p.lastDir == SOUTH {
			l = loc{p.loc.x - 1, p.loc.y}
			symbol = b[l.y][l.x]
			lastDir = EAST
		}
	} else if p.symbol == "F" {
		if p.lastDir == EAST {
			l = loc{p.loc.x, p.loc.y + 1}
			symbol = b[l.y][l.x]
			lastDir = NORTH
		}
		if p.lastDir == SOUTH {
			l = loc{p.loc.x + 1, p.loc.y}
			symbol = b[l.y][l.x]
			lastDir = WEST
		}
	} else if p.symbol == "L" {
		if p.lastDir == EAST {
			l = loc{p.loc.x, p.loc.y - 1}
			symbol = b[l.y][l.x]
			lastDir = SOUTH
		}
		if p.lastDir == NORTH {
			l = loc{p.loc.x + 1, p.loc.y}
			symbol = b[l.y][l.x]
			lastDir = WEST
		}
	} else if p.symbol == "J" {
		if p.lastDir == WEST {
			l = loc{p.loc.x, p.loc.y - 1}
			symbol = b[l.y][l.x]
			lastDir = SOUTH
		}
		if p.lastDir == NORTH {
			l = loc{p.loc.x - 1, p.loc.y}
			symbol = b[l.y][l.x]
			lastDir = EAST
		}
	}
	loop[l.y][l.x] = true
	return pipe{symbol, l, lastDir}
}

func findStartPipes(b board, s loc) []pipe {
	startPipes := make([]pipe, 0)
	if b[s.y][s.x-1] == "-" || b[s.y][s.x-1] == "L" || b[s.y][s.x-1] == "F" || b[s.y][s.x-1] == "S" {
		startPipes = append(startPipes, pipe{b[s.y][s.x-1], loc{s.x - 1, s.y}, EAST})
		loop[s.y][s.x-1] = true
	}
	if b[s.y][s.x+1] == "-" || b[s.y][s.x+1] == "7" || b[s.y][s.x+1] == "J" || b[s.y][s.x+1] == "S" {
		startPipes = append(startPipes, pipe{b[s.y][s.x+1], loc{s.x + 1, s.y}, WEST})
		loop[s.y][s.x+1] = true
	}
	if b[s.y-1][s.x] == "|" || b[s.y-1][s.x] == "7" || b[s.y-1][s.x] == "F" || b[s.y-1][s.x] == "S" {
		startPipes = append(startPipes, pipe{b[s.y-1][s.x], loc{s.x, s.y - 1}, SOUTH})
		loop[s.y-1][s.x] = true
	}
	if b[s.y+1][s.x] == "|" || b[s.y+1][s.x] == "J" || b[s.y+1][s.x] == "L" || b[s.y+1][s.x] == "S" {
		startPipes = append(startPipes, pipe{b[s.y+1][s.x], loc{s.x, s.y + 1}, NORTH})
		loop[s.y+1][s.x] = true
	}
	return startPipes
}

func getBoard(input string) [][]string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	b := make([][]string, 0)
	padLenX := len(lines[0]) + 2
	padLine := make([]string, 0)
	padLoopRow := make([]bool, 0)
	for i := 0; i < padLenX; i++ {
		padLine = append(padLine, ".")
		padLoopRow = append(padLoopRow, false)
	}
	b = append(b, padLine)
	loop = append(loop, padLoopRow)
	for _, line := range lines {
		l := make([]string, 0)
		lb := make([]bool, 0)
		l = append(l, ".")
		lb = append(lb, false)
		for _, r := range line {
			l = append(l, string(r))
			lb = append(lb, false)
		}
		l = append(l, ".")
		lb = append(lb, false)
		b = append(b, l)
		loop = append(loop, lb)
	}
	b = append(b, padLine)
	loop = append(loop, padLoopRow)

	return b
}

func findStart(board board) loc {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == "S" {
				loop[y][x] = true
				return loc{x, y}
			}
		}
	}

	return loc{-1, -1}
}
