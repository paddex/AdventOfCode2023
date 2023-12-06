package p1

import (
	"os"
	"reflect"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type test struct {
	name  string
	input race
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
	want := 288

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestGetRaces(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	got := getRaces(string(input))

	want := []race{{7, 9}, {15, 40}, {30, 200}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ergebnis: %v, Erwartet: %v", got, want)
	}
}

func TestGetNumOfPossibilities(t *testing.T) {
	tests := []test{
		{"Race 1", race{7, 9}, 4},
		{"Race 2", race{15, 40}, 8},
		{"Race 3", race{30, 200}, 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getNumOfPossibilities(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
