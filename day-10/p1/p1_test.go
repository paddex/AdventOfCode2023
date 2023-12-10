package p1

import (
	"os"
	"reflect"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type p1test struct {
	name  string
	input types.App
	want  int
}

type startTest struct {
	name  string
	input board
	want  loc
}

type startPipeTest struct {
	name   string
	inputB board
	inputS loc
	want   []pipe
}

type nextPipeTest struct {
	name   string
	inputB board
	inputP pipe
	want   pipe
}

func TestP1(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}
	input2, err := os.ReadFile("../testinput1-2")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}
	app2 := types.App{
		Logger: logger,
		Input:  string(input2),
	}

	tests := []p1test{
		{"Test 1", app, 4},
		{"Test 2", app2, 8},
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

func TestFindStart(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}
	input2, err := os.ReadFile("../testinput1-2")
	if err != nil {
		panic("Can't read file")
	}

	b1 := getBoard(string(input))
	b2 := getBoard(string(input2))

	tests := []startTest{
		{"Board 1:", b1, loc{2, 2}},
		{"Board 2:", b2, loc{1, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findStart(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestFindStartPipes(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}
	input2, err := os.ReadFile("../testinput1-2")
	if err != nil {
		panic("Can't read file")
	}

	b1 := getBoard(string(input))
	b2 := getBoard(string(input2))
	s1 := findStart(b1)
	s2 := findStart(b2)

	tests := []startPipeTest{
		{"Board 1:", b1, s1, []pipe{{"-", loc{3, 2}, WEST}, {"|", loc{2, 3}, NORTH}}},
		{"Board 2:", b2, s2, []pipe{{"J", loc{2, 3}, WEST}, {"|", loc{1, 4}, NORTH}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findStartPipes(test.inputB, test.inputS)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestFindNextPipe(t *testing.T) {
	input, err := os.ReadFile("../testinput1-2")
	if err != nil {
		panic("Can't read file")
	}

	b := getBoard(string(input))

	tests := []nextPipeTest{
		{"PIPE -", b, pipe{"-", loc{3, 4}, EAST}, pipe{"F", loc{2, 4}, EAST}},
		{"PIPE |", b, pipe{"|", loc{1, 4}, SOUTH}, pipe{"S", loc{1, 3}, SOUTH}},
		{"PIPE 7", b, pipe{"7", loc{5, 3}, WEST}, pipe{"J", loc{5, 4}, NORTH}},
		{"PIPE F", b, pipe{"F", loc{2, 2}, SOUTH}, pipe{"J", loc{3, 2}, WEST}},
		{"PIPE J", b, pipe{"J", loc{2, 3}, WEST}, pipe{"F", loc{2, 2}, SOUTH}},
		{"PIPE L", b, pipe{"L", loc{1, 5}, EAST}, pipe{"|", loc{1, 4}, SOUTH}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findNextPipe(test.inputB, test.inputP)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
