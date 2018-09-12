package cards

import "testing"

func TestNewDeck(t *testing.T) {
	newDeck := NewDeck()
	if len(newDeck.Cards) != 52 {
		t.Errorf("Wrong number of cards, got: %d, want: %d", len(newDeck.Cards), 52)
	}
	counters := make(map[string]int)
	for _, card := range newDeck.Cards {
		if _, present := counters[card.Rank]; present {
			counters[card.Rank]++
		} else {
			counters[card.Rank] = 1
		}
		if _, present := counters[card.Suit]; present {
			counters[card.Suit]++
		} else {
			counters[card.Suit] = 1
		}
	}
	tables := []struct {
		x string
		y int
	}{
		{"Spades", 13},
		{"Diamonds", 13},
		{"Hearts", 13},
		{"Clubs", 13},
		{"Two", 4},
		{"Three", 4},
		{"Four", 4},
		{"Five", 4},
		{"Six", 4},
		{"Seven", 4},
		{"Eight", 4},
		{"Nine", 4},
		{"Ten", 4},
		{"Jack", 4},
		{"Queen", 4},
		{"King", 4},
		{"Ace", 4},
	}
	for _, table := range tables {
		if counters[table.x] != table.y {
			t.Errorf("Wrong number of cards for %s, got: %d, want: %d", table.x, counters[table.x], table.y)
		}
	}

}
