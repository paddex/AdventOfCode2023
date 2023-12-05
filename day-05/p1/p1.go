package p1

import (
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type mapping struct {
	minVal       int
	maxVal       int
	mappingStart int
}

func P1(app types.App) int {
	parts := strings.Split(app.Input, "\n\n")

	seeds := getSeeds(parts[0])
	seedToSoil := getMap(parts[1])
	soilToFertilizer := getMap(parts[2])
	fertilizerToWater := getMap(parts[3])
	waterToLight := getMap(parts[4])
	lightToTemperature := getMap(parts[5])
	temperatureToHumidity := getMap(parts[6])
	humidityToLocation := getMap(parts[7])

	closestLoc := -1
	for _, seed := range seeds {
		soil := getNext(seed, seedToSoil)
		fert := getNext(soil, soilToFertilizer)
		water := getNext(fert, fertilizerToWater)
		light := getNext(water, waterToLight)
		temp := getNext(light, lightToTemperature)
		humidity := getNext(temp, temperatureToHumidity)
		location := getNext(humidity, humidityToLocation)

		if location < closestLoc || closestLoc < 0 {
			closestLoc = location
		}
	}
	return closestLoc
}

func getSeeds(seedStr string) []int {
	seedValsStr := strings.TrimSpace(strings.Split(seedStr, ":")[1])
	myseeds := strings.Split(seedValsStr, " ")
	seeds := make([]int, 0)
	for _, seed := range myseeds {
		seedI, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedI)
	}

	return seeds
}

func getMap(mapStr string) []mapping {
	mapLines := strings.Split(mapStr, "\n")[1:]
	resMap := make([]mapping, 0)
	for _, line := range mapLines {
		if line == "" {
			continue
		}
		resMapping := mapLine(line)
		resMap = append(resMap, resMapping)
	}

	return resMap
}

func mapLine(line string) mapping {
	parts := strings.Split(line, " ")
	toMap, _ := strconv.Atoi(parts[0])
	val, _ := strconv.Atoi(parts[1])
	length, _ := strconv.Atoi(parts[2])

	return mapping{minVal: val, maxVal: val + length, mappingStart: toMap}
}

func getNext(val int, valMap []mapping) int {
	for _, valMapping := range valMap {
		if val >= valMapping.minVal && val <= valMapping.maxVal {
			diff := val - valMapping.minVal
			return valMapping.mappingStart + diff
		}
	}
	return val
}
