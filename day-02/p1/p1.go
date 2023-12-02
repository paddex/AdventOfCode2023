package p1

import (
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

var maxcolors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func P1(app types.App) int {
	valids := make([]int, 0)

	lines := strings.Split(app.Input, "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		valid := true
		colormap := processLine(line, app)
		if colormap["red"] > maxcolors["red"] {
			valid = false
		}
		if colormap["green"] > maxcolors["green"] {
			valid = false
		}
		if colormap["blue"] > maxcolors["blue"] {
			valid = false
		}

		if valid {
			valids = append(valids, colormap["game"])
		}
	}

	sum := 0
	for _, valid := range valids {
		sum += valid
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
