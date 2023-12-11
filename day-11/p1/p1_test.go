package p1

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type appTest struct {
	name  string
	input types.App
	want  int
}

func TestP1(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	tests := []appTest{
		{"testinput1", app, 374},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := P1(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

type funcTest struct {
	name  string
	input pair
	gMap  map[int]loc
	want  int
}

func TestGetDistance(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	universe := getUniverse(string(input))
	universe = expandUniverse(universe)
	universe, galaxyCount, galaxyMap := numberGalaxies(universe)

	_ = galaxyCount

	tests := []funcTest{
		{"5-9", pair{5, 9}, galaxyMap, 9},
		{"9-5", pair{9, 5}, galaxyMap, 9},
		{"1-7", pair{1, 7}, galaxyMap, 15},
		{"3-6", pair{3, 6}, galaxyMap, 17},
		{"8-9", pair{8, 9}, galaxyMap, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getDistance(test.input, test.gMap)
			_ = got

			// if got != [][]string{} {
			// 	t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			// }
		})
	}
}
