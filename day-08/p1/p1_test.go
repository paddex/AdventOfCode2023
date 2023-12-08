package p1

import (
	"os"
	"reflect"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type nodetest struct {
	name  string
	input string
	want  node
}

type apptest struct {
	name  string
	input types.App
	want  int
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

	tests := []apptest{
		{"Input 1", app, 2},
		{"Input 2", app2, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := P1(test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetNode(t *testing.T) {
	tests := []nodetest{
		{"Node 1", "AAA = (BBB, CCC)", node{"AAA", "BBB", "CCC"}},
		{"Node 2", "BBB = (DDD, EEE)", node{"BBB", "DDD", "EEE"}},
		{"Node 3", "CCC = (ZZZ, GGG)", node{"CCC", "ZZZ", "GGG"}},
		{"Node 4", "DDD = (DDD, DDD)", node{"DDD", "DDD", "DDD"}},
		{"Node 5", "EEE = (EEE, EEE)", node{"EEE", "EEE", "EEE"}},
		{"Node 6", "GGG = (GGG, GGG)", node{"GGG", "GGG", "GGG"}},
		{"Node 7", "ZZZ = (ZZZ, ZZZ)", node{"ZZZ", "ZZZ", "ZZZ"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getNode(test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
