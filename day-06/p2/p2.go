package p2

import (
	"fmt"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type race struct {
	time     int
	distance int
}

func P2(app types.App) int {
	race := getRace(app.Input)

	res := getNumOfPossibilities(race)

	return res
}

func getRace(input string) race {
	lines := strings.Split(input, "\n")
	timeLine := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	distLine := strings.TrimSpace(strings.Split(lines[1], ":")[1])

	times := strings.Fields(timeLine)
	dists := strings.Fields(distLine)

	time, _ := strconv.Atoi(strings.Join(times, ""))
	dist, _ := strconv.Atoi(strings.Join(dists, ""))

	return race{time, dist}
}

func getNumOfPossibilities(race race) int {
	for i := 0; i <= race.time/2; i++ {
		if i*(race.time-i) > race.distance {
			fmt.Printf("%d\n", i)
			return (race.time + 1) - 2*i
		}
	}
	return 0
}
