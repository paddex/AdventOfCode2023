package p1

import (
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

func P1(app types.App) int {
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

	return (steps / 2) + 1
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
	return pipe{symbol, l, lastDir}
}

func findStartPipes(b board, s loc) []pipe {
	startPipes := make([]pipe, 0)
	if b[s.y][s.x-1] == "-" || b[s.y][s.x-1] == "L" || b[s.y][s.x-1] == "F" || b[s.y][s.x-1] == "S" {
		startPipes = append(startPipes, pipe{b[s.y][s.x-1], loc{s.x - 1, s.y}, EAST})
	}
	if b[s.y][s.x+1] == "-" || b[s.y][s.x+1] == "7" || b[s.y][s.x+1] == "J" || b[s.y][s.x+1] == "S" {
		startPipes = append(startPipes, pipe{b[s.y][s.x+1], loc{s.x + 1, s.y}, WEST})
	}
	if b[s.y-1][s.x] == "|" || b[s.y-1][s.x] == "7" || b[s.y-1][s.x] == "F" || b[s.y-1][s.x] == "S" {
		startPipes = append(startPipes, pipe{b[s.y-1][s.x], loc{s.x, s.y - 1}, SOUTH})
	}
	if b[s.y+1][s.x] == "|" || b[s.y+1][s.x] == "J" || b[s.y+1][s.x] == "L" || b[s.y+1][s.x] == "S" {
		startPipes = append(startPipes, pipe{b[s.y+1][s.x], loc{s.x, s.y + 1}, NORTH})
	}
	return startPipes
}

func getBoard(input string) [][]string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	b := make([][]string, 0)
	padLenX := len(lines[0]) + 2
	padLine := make([]string, 0)
	for i := 0; i < padLenX; i++ {
		padLine = append(padLine, ".")
	}
	b = append(b, padLine)
	for _, line := range lines {
		l := make([]string, 0)
		l = append(l, ".")
		for _, r := range line {
			l = append(l, string(r))
		}
		l = append(l, ".")
		b = append(b, l)
	}

	return b
}

func findStart(board board) loc {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == "S" {
				return loc{x, y}
			}
		}
	}

	return loc{-1, -1}
}
