package p1

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

func P1(app types.App) int {
	races := getRaces(app.Input)

	product := 1
	for _, race := range races {
		possibilities := getNumOfPossibilities(race)
		product *= possibilities
	}

	return product
}

func getRaces(input string) []race {
	lines := strings.Split(input, "\n")
	timeLine := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	distLine := strings.TrimSpace(strings.Split(lines[1], ":")[1])

	times := strings.Fields(timeLine)
	dists := strings.Fields(distLine)

	races := make([]race, 0)
	for i, time := range times {
		timeVal, _ := strconv.Atoi(time)
		distVal, _ := strconv.Atoi(dists[i])
		races = append(races, race{timeVal, distVal})
	}
	return races
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
