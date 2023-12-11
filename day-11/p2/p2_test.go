package p2

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

func TestP2(t *testing.T) {
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
		{"testinput1", app, 1030},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := P2(test.input)

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
		{"1-2", pair{1, 2}, galaxyMap, 14},
		{"2-1", pair{2, 1}, galaxyMap, 14},
		{"2-3", pair{2, 3}, galaxyMap, 26},
		{"2-5", pair{2, 5}, galaxyMap, 37},
		{"5-2", pair{5, 2}, galaxyMap, 37},
		{"2-9", pair{2, 9}, galaxyMap, 38},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getDistance(test.input, test.gMap, universe)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
