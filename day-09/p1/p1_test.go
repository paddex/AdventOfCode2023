package p1

import (
	"os"
	"reflect"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type valTest struct {
	name  string
	input []int
	want  int
}

type arrTest struct {
	name  string
	input string
	want  []int
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
	want := 114

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestGetIntArray(t *testing.T) {
	tests := []arrTest{
		{"Row 1", "0 3 6 9 12 15", []int{0, 3, 6, 9, 12, 15}},
		{"Row 2", "1 3 6 10 15 21", []int{1, 3, 6, 10, 15, 21}},
		{"Row 3", "10 13 16 21 30 45", []int{10, 13, 16, 21, 30, 45}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getIntArr(test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetNextValue(t *testing.T) {
	tests := []valTest{
		{"Row 1", []int{0, 3, 6, 9, 12, 15}, 18},
		{"Row 2", []int{1, 3, 6, 10, 15, 21}, 28},
		{"Row 3", []int{10, 13, 16, 21, 30, 45}, 68},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getNextValue(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
