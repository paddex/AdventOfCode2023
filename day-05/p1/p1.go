package p1

import (
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

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
	_ = seeds
	_ = seedToSoil
	_ = soilToFertilizer
	_ = fertilizerToWater
	_ = waterToLight
	_ = lightToTemperature
	_ = temperatureToHumidity
	_ = humidityToLocation

	closestLoc := 9999999
	for _, seed := range seeds {
		soil := getSoil(seed, seedToSoil)
		fert := getFertilizer(soil, soilToFertilizer)
		water := getWater(fert, fertilizerToWater)
		light := getLight(water, waterToLight)
		temp := getTemperature(light, lightToTemperature)
		humidity := getHumidity(temp, temperatureToHumidity)
		location := getLocation(humidity, humidityToLocation)

		if location < closestLoc {
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

func getMap(mapStr string) map[int]int {
	mapLines := strings.Split(mapStr, "\n")[1:]
	resMap := make(map[int]int)
	for _, line := range mapLines {
		mapLine(line, resMap)
	}

	return resMap
}

func mapLine(line string, resMap map[int]int) {
	if line == "" {
		return
	}
	parts := strings.Split(line, " ")
	toMap, _ := strconv.Atoi(parts[0])
	val, _ := strconv.Atoi(parts[1])
	length, _ := strconv.Atoi(parts[2])

	for i := 0; i < length; i++ {
		resMap[val+i] = toMap + i
	}
}

func getSoil(seed int, seedToSoil map[int]int) int {
	if soil, ok := seedToSoil[seed]; ok {
		return soil
	}
	return seed
}

func getFertilizer(soil int, soilToFertilizer map[int]int) int {
	if fert, ok := soilToFertilizer[soil]; ok {
		return fert
	}
	return soil
}

func getWater(fertilizer int, fertilizerToWater map[int]int) int {
	if water, ok := fertilizerToWater[fertilizer]; ok {
		return water
	}
	return fertilizer
}

func getLight(water int, waterToLight map[int]int) int {
	if light, ok := waterToLight[water]; ok {
		return light
	}
	return water
}

func getTemperature(light int, lightToTemperature map[int]int) int {
	if temp, ok := lightToTemperature[light]; ok {
		return temp
	}
	return light
}

func getHumidity(temperature int, temperatureToHumidity map[int]int) int {
	if humidity, ok := temperatureToHumidity[temperature]; ok {
		return humidity
	}
	return temperature
}

func getLocation(humidity int, humidityToLocation map[int]int) int {
	if location, ok := humidityToLocation[humidity]; ok {
		return location
	}
	return humidity
}
