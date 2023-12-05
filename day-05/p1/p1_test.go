package p1

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type test struct {
	name  string
	input int
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

	got := P1(app)
	want := 35

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
