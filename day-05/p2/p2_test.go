package p2

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type test struct {
	name  string
	input int64
	want  int64
}

func TestP2(t *testing.T) {
	input, err := os.ReadFile("../testinput2")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	got := P2(app)
	var want int64 = 46

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestGetNext(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	seedToSoil := getMap(parts[1])

	tests := []test{
		{"Seed 79", 79, 81},
		{"Seed 14", 14, 14},
		{"Seed 55", 55, 57},
		{"Seed 13", 13, 13},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getNext(test.input, seedToSoil)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetSeeds(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	seedRanges := getSeeds(parts[0])

	want := []seedRange{{start: 79, length: 14}, {start: 55, length: 13}}

	if !reflect.DeepEqual(seedRanges, want) {
		t.Errorf("Ergebnis: %v, Erwartet: %v", seedRanges, want)
	}
}

func TestMyLoop(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	seedRanges := getSeeds(parts[0])

	want := []int64{92, 67}
	got := make([]int64, 0)
	var res int64 = 0
	for _, seedRange := range seedRanges {
		res = 0
		for i := seedRange.start; i < seedRange.start+seedRange.length; i++ {
			res = i
		}
		got = append(got, res)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ergebnis: %v, Erwartet: %v", got, want)
	}
}
