package p2

import (
	"fmt"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type mapping struct {
	minVal       int64
	maxVal       int64
	mappingStart int64
}

type seedRange struct {
	start  int64
	length int64
}

func P2(app types.App) int64 {
	parts := strings.Split(app.Input, "\n\n")

	seedRanges := getSeeds(parts[0])
	seedToSoil := getMap(parts[1])
	soilToFertilizer := getMap(parts[2])
	fertilizerToWater := getMap(parts[3])
	waterToLight := getMap(parts[4])
	lightToTemperature := getMap(parts[5])
	temperatureToHumidity := getMap(parts[6])
	humidityToLocation := getMap(parts[7])

	ch := make(chan int64)
	for _, sr := range seedRanges {
		go func(nsr seedRange) {
			var closestLoc int64 = -1
			for i := nsr.start; i < nsr.start+nsr.length; i++ {
				soil := getNext(i, seedToSoil)
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
			ch <- closestLoc
		}(sr)
	}

	allClosest := make([]int64, 0)
	for range seedRanges {
		closest := <-ch
		fmt.Printf("GOT ONE: %d\n", closest)
		allClosest = append(allClosest, closest)
	}

	var realClosest int64 = -1
	for _, cl := range allClosest {
		if cl < realClosest || realClosest < 0 {
			realClosest = cl
		}
	}

	return realClosest
}

func getSeeds(seedStr string) []seedRange {
	seedValsStr := strings.TrimSpace(strings.Split(seedStr, ":")[1])
	myseeds := strings.Split(seedValsStr, " ")
	seeds := make([]seedRange, 0)
	rng := make([]int64, 0)
	for _, seed := range myseeds {
		seedI, _ := strconv.ParseInt(seed, 10, 64)
		rng = append(rng, seedI)
		if len(rng) == 2 {
			seeds = append(seeds, seedRange{start: rng[0], length: rng[1]})
			rng = make([]int64, 0)
		}
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
	toMap, _ := strconv.ParseInt(parts[0], 10, 64)
	val, _ := strconv.ParseInt(parts[1], 10, 64)
	length, _ := strconv.ParseInt(parts[2], 10, 64)

	return mapping{minVal: val, maxVal: val + length, mappingStart: toMap}
}

func getNext(val int64, valMap []mapping) int64 {
	for _, valMapping := range valMap {
		if val >= valMapping.minVal && val < valMapping.maxVal {
			diff := val - valMapping.minVal
			return valMapping.mappingStart + diff
		}
	}
	return val
}
