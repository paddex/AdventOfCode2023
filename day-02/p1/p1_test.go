package p1

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

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
	want := 8

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestProcessLine(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	type test struct {
		name  string
		input string
		want  map[string]int
	}

	lines := strings.Split(string(input), "\n")

	tests := []test{
		{"Game 1", lines[0], map[string]int{"blue": 6, "red": 4, "green": 2, "game": 1}},
		{"Game 2", lines[1], map[string]int{"blue": 4, "red": 1, "green": 3, "game": 2}},
		{"Game 3", lines[2], map[string]int{"blue": 6, "red": 20, "green": 13, "game": 3}},
		{"Game 4", lines[3], map[string]int{"blue": 15, "red": 14, "green": 3, "game": 4}},
		{"Game 5", lines[4], map[string]int{"blue": 2, "red": 6, "green": 3, "game": 5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLine(test.input, app)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
