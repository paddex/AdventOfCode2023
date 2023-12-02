package p2

import (
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

func P2(app types.App) int {
	powers := make([]int, 0)

	lines := strings.Split(app.Input, "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		colormap := processLine(line, app)
		power := colormap["red"] * colormap["green"] * colormap["blue"]
		powers = append(powers, power)
	}

	sum := 0
	for _, power := range powers {
		sum += power
	}

	return sum
}

func processLine(line string, app types.App) map[string]int {
	colormap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
		"game":  0,
	}

	lineparts := strings.Split(line, ":")
	gameid, err := strconv.Atoi(strings.Split(lineparts[0], " ")[1])
	if err != nil {
		app.Logger.Error(err.Error())
		panic(err)
	}
	colormap["game"] = gameid
	gamepart := strings.Trim(lineparts[1], " ")
	games := strings.Split(gamepart, ";")

	for _, game := range games {
		colors := strings.Split(game, ",")
		for _, color := range colors {
			color = strings.Trim(color, " ")
			split := strings.Split(color, " ")
			count, err := strconv.Atoi(split[0])
			if err != nil {
				app.Logger.Error(err.Error())
				panic(err)
			}
			colorname := split[1]

			if count > colormap[colorname] {
				colormap[colorname] = count
			}
		}
	}
	return colormap
}
