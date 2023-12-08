package p2

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

func P2(app types.App) int {
	lines := strings.Split(app.Input, "\n")
	lr := lines[0]
	lines = lines[2 : len(lines)-1]

	nodes := make(map[string]node)
	for _, line := range lines {
		n := getNode(line)
		nodes[n.name] = n
	}

	startingNodes := make([]string, 0)
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			startingNodes = append(startingNodes, k)
		}
	}

	dirs := getDirs(lr)
	steps := 0
	stepAmounts := make([]int, 0)
	finished := false
	currentDirIdx := 0
	for !finished {
		finished = true
		dir := dirs[currentDirIdx]
		steps++
		for i, n := range startingNodes {
			if !(strings.HasSuffix(n, "Z")) {
				finished = false
				var nextNode string

				if dir == LEFT {
					nextNode = nodes[n].left
				} else {
					nextNode = nodes[n].right
				}

				startingNodes[i] = nextNode

				if strings.HasSuffix(nextNode, "Z") {
					stepAmounts = append(stepAmounts, steps)
				}
			}
		}
		currentDirIdx = getNextDirIndex(dirs, currentDirIdx)
	}

	lcm := LCM(stepAmounts[0], stepAmounts[1], stepAmounts[2:]...)

	return lcm
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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
