package p1

import (
	"strings"

	"paddex.net/aoc/types"
)

type node struct {
	name  string
	left  string
	right string
}

type dir int

const (
	LEFT dir = iota
	RIGHT
)

func P1(app types.App) int {
	lines := strings.Split(app.Input, "\n")
	lr := lines[0]
	lines = lines[2 : len(lines)-1]

	nodes := make(map[string]node)
	for _, line := range lines {
		n := getNode(line)
		nodes[n.name] = n
	}

	dirs := getDirs(lr)
	currentDirIdx := 0
	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		dir := dirs[currentDirIdx]
		var nextNode string

		if dir == LEFT {
			nextNode = nodes[currentNode].left
		} else {
			nextNode = nodes[currentNode].right
		}

		currentDirIdx = getNextDirIndex(dirs, currentDirIdx)
		currentNode = nextNode
		steps++
	}

	return steps
}

func getNode(input string) node {
	parts := strings.Split(input, " = ")
	name := parts[0]
	nodeParts := strings.Split(parts[1], ", ")
	left := strings.TrimLeft(nodeParts[0], "(")
	right := strings.TrimRight(nodeParts[1], ")")
	return node{name, left, right}
}

func getDirs(input string) []dir {
	dirs := make([]dir, 0)
	for _, r := range input {
		if string(r) == "L" {
			dirs = append(dirs, LEFT)
		} else if string(r) == "R" {
			dirs = append(dirs, RIGHT)
		}
	}

	return dirs
}

func getNextDirIndex(directions []dir, current int) int {
	if current >= len(directions)-1 {
		return 0
	}

	current++

	return current
}
