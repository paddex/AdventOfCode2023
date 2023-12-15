package p2

import (
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type box struct {
	lenses []lens
}

type lens struct {
	label string
	flen  int
}

func P2(app types.App) int {
	fields := strings.Split(strings.TrimSpace(app.Input), ",")

	power := 0

	boxes := make([]box, 255)
	for _, field := range fields {
		label, op, flen := getLabelName(field)
		box := hashString(label)
		if op == "=" {
			if len(boxes[box].lenses) == 0 {
				boxes[box].lenses = append(boxes[box].lenses, lens{label, flen})
			} else {
				lensIndex := containsLabel(boxes[box].lenses, label)
				if lensIndex >= 0 {
					boxes[box].lenses[lensIndex] = lens{label, flen}
				} else {
					boxes[box].lenses = append(boxes[box].lenses, lens{label, flen})
				}
			}
		} else if op == "-" {
			boxes[box].lenses = filterBox(boxes[box].lenses, label)
		}
	}

	/*
		for _, box := range boxes {
			if len(box.lenses) != 0 {
				fmt.Printf("%v\n", box)
			}
		}
	*/

	for i, box := range boxes {
		for j, l := range box.lenses {
			power += focusPower(i, j, l)
		}
	}
	return power
}

func focusPower(boxnum int, lensnum int, lens lens) int {
	return (boxnum + 1) * (lensnum + 1) * lens.flen
}

func filterBox(lenses []lens, label string) []lens {
	nlenses := make([]lens, 0)
	for _, l := range lenses {
		if l.label != label {
			nlenses = append(nlenses, l)
		}
	}
	return nlenses
}

func containsLabel(lenses []lens, label string) int {
	for i, l := range lenses {
		if l.label == label {
			return i
		}
	}
	return -1
}

func getLabelName(input string) (string, string, int) {
	idx := strings.Index(input, "=")
	if idx >= 0 {
		flen, _ := strconv.Atoi(input[idx+1:])
		return input[:idx], "=", flen
	}
	idx = strings.Index(input, "-")
	if idx > 0 {
		return input[:idx], "-", 0
	}

	return input, "", 0
}

func hashString(input string) int {
	hash := 0
	for i := 0; i < len(input); i++ {
		c := int(input[i])
		hash += c
		hash *= 17
		hash %= 256
	}
	return hash
}
