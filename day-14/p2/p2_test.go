package p2

import (
	"os"
	"reflect"
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
		{"testinput1", app, 64},
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

type rotateTest struct {
	name  string
	input [][]string
	want  [][]string
}

func TestRotate(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	inp := getGrid(string(input))

	resC1 := `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`

	resC2 := `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`

	resC3 := `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`

	c1 := getGrid(resC1)
	c2 := getGrid(resC2)
	c3 := getGrid(resC3)

	tests := []rotateTest{
		{"cycle 1", inp, c1},
		{"cycle 2", c1, c2},
		{"cycle 3", c2, c3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := rotate(test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, "\n"+getString(got), "\n"+getString(test.want))
			}
		})
	}
}

type cycleTest struct {
	name  string
	input []string
	lam   int
	mu    int
}

func TestFindCycle(t *testing.T) {
	t1 := make([]string, 0)
	t1 = append(t1, "1")
	t1 = append(t1, "2")
	for i := 0; i < 5; i++ {
		t1 = append(t1, "3")
		t1 = append(t1, "4")
		t1 = append(t1, "5")
	}
	t1 = append(t1, "3")
	tests := []cycleTest{
		{"test 1", t1, 3, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lam, mu := findCycle(test.input)

			if lam != test.lam && mu != test.mu {
				t.Errorf("TEST: %s: got lam %v, want lam %v, got mu %v, want mu %v", test.name, lam, test.lam, mu, test.mu)
			}
		})
	}
}
