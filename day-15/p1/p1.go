package p1

import (
	"strings"

	"paddex.net/aoc/types"
)

func P1(app types.App) int {
	fields := strings.Split(strings.TrimSpace(app.Input), ",")

	sum := 0
	for _, field := range fields {
		sum += hashString(field)
	}
	return sum
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
