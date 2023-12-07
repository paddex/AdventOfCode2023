package p2

import (
	"os"
	"reflect"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type test struct {
	name  string
	input string
	want  hand
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
	// want := 5905
	want := 4657

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestGetHand(t *testing.T) {
	tests := []test{
		{"Hand 1", "32T3K 765", hand{[]cardtype{THREE, TWO, TEN, THREE, KING}, ONEPAIR, 765}},
		{"Hand 2", "T55J5 684", hand{[]cardtype{TEN, FIVE, FIVE, JACK, FIVE}, FOURKIND, 684}},
		{"Hand 3", "KK677 28", hand{[]cardtype{KING, KING, SIX, SEVEN, SEVEN}, TWOPAIR, 28}},
		{"Hand 4", "KTJJT 220", hand{[]cardtype{KING, TEN, JACK, JACK, TEN}, FOURKIND, 220}},
		{"Hand 5", "QQQJA 483", hand{[]cardtype{QUEEN, QUEEN, QUEEN, JACK, ACE}, FOURKIND, 483}},
		{"Hand 6", "55555 483", hand{[]cardtype{FIVE, FIVE, FIVE, FIVE, FIVE}, FIVEKIND, 483}},
		{"Hand 7", "44442 483", hand{[]cardtype{FOUR, FOUR, FOUR, FOUR, TWO}, FOURKIND, 483}},
		{"Hand 8", "33322 483", hand{[]cardtype{THREE, THREE, THREE, TWO, TWO}, FULLHOUSE, 483}},
		{"Hand 9", "JJJJ2 483", hand{[]cardtype{JACK, JACK, JACK, JACK, TWO}, FIVEKIND, 483}},
		{"Hand 10", "J2233 483", hand{[]cardtype{JACK, TWO, TWO, THREE, THREE}, FULLHOUSE, 483}},
		{"Hand 11", "J2435 483", hand{[]cardtype{JACK, TWO, FOUR, THREE, FIVE}, ONEPAIR, 483}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseHand(test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
