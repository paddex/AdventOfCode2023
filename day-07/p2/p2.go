package p2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"paddex.net/aoc/types"
)

type (
	handtype int
	cardtype int
)

const (
	HI handtype = iota
	ONEPAIR
	TWOPAIR
	THREEKIND
	FULLHOUSE
	FOURKIND
	FIVEKIND
)

const (
	JACK cardtype = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	QUEEN
	KING
	ACE
)

type hand struct {
	cards    []cardtype
	handtype handtype
	bid      int
}

func P2(app types.App) int {
	lines := strings.Split(app.Input, "\n")
	lines = lines[:len(lines)-1]

	hands := make([]hand, 0)
	for _, line := range lines {
		h := parseHand(line)
		hands = append(hands, h)
	}

	hands = sortHands(hands)

	fmt.Printf("%v\n", hands)

	sum := 0
	for i, h := range hands {
		val := (i + 1) * h.bid
		sum += val
	}

	return sum
}

func sortHands(hands []hand) []hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handtype == hands[j].handtype {
			for k := 0; k < 5; k++ {
				if !(hands[i].cards[k] == hands[j].cards[k]) {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
		}

		return hands[i].handtype < hands[j].handtype
	})

	return hands
}

func parseHand(handStr string) hand {
	parts := strings.Split(handStr, " ")
	bid, _ := strconv.Atoi(parts[1])

	charOccurrences := make(map[string]int)
	handCards := make([]cardtype, 0)

	for _, char := range parts[0] {
		if _, exists := charOccurrences[string(char)]; exists {
			charOccurrences[string(char)]++
		} else {
			charOccurrences[string(char)] = 1
		}
		handCards = append(handCards, getCardType(string(char)))
	}

	var maxChar string
	maxOcc := 0
	for k, v := range charOccurrences {
		if k == "J" {
			continue
		}
		if v > maxOcc {
			maxChar = k
			maxOcc = v
		}
	}

	var handtype handtype
	numJacks := charOccurrences["J"]
	have3 := false
	num2 := 0
	charOccurrences[maxChar] += numJacks
	for k, v := range charOccurrences {
		if v < 2 || k == "J" {
			continue
		}
		if v == 5 {
			handtype = FIVEKIND
			break
		}
		if v == 4 {
			handtype = FOURKIND
			break
		}
		if v == 3 {
			have3 = true
		}
		if v == 2 {
			num2++
		}
	}

	if numJacks == 5 {
		handtype = FIVEKIND
	}

	if have3 {
		if num2 > 0 {
			handtype = FULLHOUSE
		} else {
			handtype = THREEKIND
		}
	} else if num2 > 1 {
		handtype = TWOPAIR
	} else if num2 == 1 {
		handtype = ONEPAIR
	}

	return hand{cards: handCards, handtype: handtype, bid: bid}
}

func getCardType(char string) cardtype {
	switch char {
	case "2":
		return TWO
	case "3":
		return THREE
	case "4":
		return FOUR
	case "5":
		return FIVE
	case "6":
		return SIX
	case "7":
		return SEVEN
	case "8":
		return EIGHT
	case "9":
		return NINE
	case "T":
		return TEN
	case "J":
		return JACK
	case "Q":
		return QUEEN
	case "K":
		return KING
	case "A":
		return ACE
	}

	return -1
}
