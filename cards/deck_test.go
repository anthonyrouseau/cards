package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStandardDeck(t *testing.T) {
	assert := assert.New(t)
	t.Run("no jokers", func(t *testing.T) {
		deck := NewStandardDeck(false)
		assert.Len(deck.cards, 52)
		assert.Equal(52, deck.maxSize)
		var ranks [14]int
		suits := map[SuitName]int{}
		colors := map[SuitColor]int{}
		for _, card := range deck.cards {
			if assert.Less(int(card.rank), len(ranks)) && assert.GreaterOrEqual(int(card.rank), 1) {
				ranks[card.rank]++
			}
			suits[card.suit.name]++
			colors[card.suit.color]++
		}
		gotRanks := ranks[1:]
		for _, rank := range gotRanks {
			assert.Equal(4, ranks[rank])
		}
		for _, count := range suits {
			assert.Equal(13, count)
		}
		for _, count := range colors {
			assert.Equal(26, count)
		}
	})
	t.Run("with jokers", func(t *testing.T) {
		deck := NewStandardDeck(true)
		assert.Len(deck.cards, 54)
		assert.Equal(54, deck.maxSize)
		var ranks [16]int
		suits := map[SuitName]int{}
		colors := map[SuitColor]int{}
		for _, card := range deck.cards {
			if assert.Less(int(card.rank), len(ranks)) && assert.GreaterOrEqual(int(card.rank), 0) {
				ranks[card.rank]++
			}
			suits[card.suit.name]++
			colors[card.suit.color]++
		}
		gotRanks := ranks[1:]
		for _, rank := range gotRanks {
			if rank < 14 {
				assert.Equal(4, ranks[rank])
			} else {
				assert.Equal(1, ranks[rank])
			}
		}
		for name, count := range suits {
			if name != Joker {
				assert.Equal(13, count)
			} else {
				assert.Equal(2, count)
			}
		}
		for _, count := range colors {
			assert.Equal(27, count)
		}
	})
}

func TestShuffle(t *testing.T) {
	assert := assert.New(t)
	var failed int
	for times := 0; times < 100; times++ {
		deck := NewStandardDeck(false)
		initial := make([]Card, len(deck.cards))
		for i, card := range deck.cards {
			initial[i] = card
		}
		deck.Shuffle()
		var sameCount int
		for i, card := range initial {
			if card.rank == deck.cards[i].rank && card.suit.name == deck.cards[i].suit.name {
				sameCount++
			}
		}
		if sameCount > 10 {
			failed++
		}
	}
	assert.Less(failed, 2)
}

func Test_Deck_Pick(t *testing.T) {
	assert := assert.New(t)
	t.Run("multiple", func(t *testing.T) {
		validIndices := []int{1, 5, 10}
		expected := make([]Card, len(validIndices))
		deck := NewStandardDeck(false)
		for i, j := range validIndices {
			expected[i] = deck.cards[j]
		}
		got, err := deck.Pick(validIndices)
		if assert.NoError(err) {
			for _, card := range got {
				assert.Contains(expected, card)
			}
		}
		if assert.Len(deck.cards, deck.maxSize-len(validIndices)) {
			for _, deckCard := range deck.cards {
				for _, gotCard := range got {
					assert.NotEqual(deckCard, gotCard)
				}
			}
		}
	})
	t.Run("single", func(t *testing.T) {
		validIndices := []int{20}
		expected := make([]Card, len(validIndices))
		deck := NewStandardDeck(false)
		for i, j := range validIndices {
			expected[i] = deck.cards[j]
		}
		got, err := deck.Pick(validIndices)
		if assert.NoError(err) {
			for _, card := range got {
				assert.Contains(expected, card)
			}
		}
		if assert.Len(deck.cards, deck.maxSize-len(validIndices)) {
			for _, deckCard := range deck.cards {
				for _, gotCard := range got {
					assert.NotEqual(deckCard, gotCard)
				}
			}
		}
	})
	t.Run("invalid", func(t *testing.T) {
		deck := NewStandardDeck(false)
		invalidIndices := []int{len(deck.cards)}
		_, err := deck.Pick(invalidIndices)
		assert.Error(err)
	})
	t.Run("repeated", func(t *testing.T) {
		deck := NewStandardDeck(false)
		repeatedIndices := []int{1, 1, 1, 1}
		_, err := deck.Pick(repeatedIndices)
		assert.Error(err)
	})
}

func BenchmarkPick(b *testing.B) {
	pick := []int{0, 1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deck := NewStandardDeck(false)
		deck.Pick(pick)
	}
}

func BenchmarkPickIndividual(b *testing.B) {
	pick := []int{10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deck := NewStandardDeck(false)
		for j := 0; j < 3; j++ {
			deck.Pick(pick)
		}
	}
}

func TestDeal(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid", func(t *testing.T) {
		deck := NewStandardDeck(false)
		numHands := 5
		sizeHands := 5
		hands, err := deck.Deal(numHands, sizeHands)
		assert.NoError(err)
		if assert.Len(hands, numHands) {
			for _, hand := range hands {
				assert.Len(hand.cards, sizeHands)
				assert.Equal(sizeHands, hand.maxSize)
			}
		}
	})
	t.Run("not enough cards", func(t *testing.T) {
		deck := NewStandardDeck(false)
		_, err := deck.Deal(30, 2)
		assert.Error(err)
	})

}

func Test_Deck_CardCount(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	assert.Len(deck.cards, deck.CardCount())
}

func Test_Deck_HasCard(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("in deck", func(t *testing.T) {
		card := deck.cards[20]
		assert.True(deck.HasCard(&card))
	})
	t.Run("not in deck", func(t *testing.T) {
		card := &Card{}
		assert.False(deck.HasCard(card))
	})
}

func Test_Deck_MaxSize(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	assert.Equal(deck.maxSize, deck.MaxSize())
}

func TestPickTop(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	expected := deck.cards[0]
	card, err := deck.PickTop()
	assert.NoError(err)
	assert.Equal(expected, card)
	assert.Len(deck.cards, deck.maxSize-1)
	assert.NotEqual(expected, deck.cards[0])
}

func BenchmarkPickTop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck := NewStandardDeck(false)
		deck.PickTop()
	}
}

func TestPickBottom(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	expected := deck.cards[len(deck.cards)-1]
	card, err := deck.PickBottom()
	assert.NoError(err)
	assert.Equal(expected, card)
	assert.Len(deck.cards, deck.maxSize-1)
	assert.NotEqual(expected, deck.cards[len(deck.cards)-1])
}

func BenchmarkPickBottom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck := NewStandardDeck(false)
		deck.PickBottom()
	}
}

func Test_Deck_Peek(t *testing.T) {
	assert := assert.New(t)
	t.Run("multiple", func(t *testing.T) {
		validIndices := []int{1, 5, 10}
		expected := make([]Card, len(validIndices))
		deck := NewStandardDeck(false)
		for i, j := range validIndices {
			expected[i] = deck.cards[j]
		}
		got, err := deck.Peek(validIndices)
		if assert.NoError(err) {
			for _, card := range got {
				assert.Contains(expected, card)
			}
		}
		assert.Len(deck.cards, deck.maxSize)
	})
	t.Run("single", func(t *testing.T) {
		validIndices := []int{20}
		expected := make([]Card, len(validIndices))
		deck := NewStandardDeck(false)
		for i, j := range validIndices {
			expected[i] = deck.cards[j]
		}
		got, err := deck.Peek(validIndices)
		if assert.NoError(err) {
			for _, card := range got {
				assert.Contains(expected, card)
			}
		}
		assert.Len(deck.cards, deck.maxSize)
	})
	t.Run("invalid", func(t *testing.T) {
		deck := NewStandardDeck(false)
		invalidIndices := []int{len(deck.cards)}
		_, err := deck.Peek(invalidIndices)
		assert.Error(err)
	})
}

func TestPeekTop(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	expected := deck.cards[0]
	card, err := deck.PeekTop()
	assert.NoError(err)
	assert.Equal(expected, card)
	assert.Len(deck.cards, deck.maxSize)
	assert.Equal(expected, deck.cards[0])
}

func TestPeekBottom(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	expected := deck.cards[len(deck.cards)-1]
	card, err := deck.PeekBottom()
	assert.NoError(err)
	assert.Equal(expected, card)
	assert.Len(deck.cards, deck.maxSize)
	assert.Equal(expected, deck.cards[len(deck.cards)-1])
}

func Test_Deck_PickRandom(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	card, err := deck.PickRandom()
	assert.NoError(err)
	assert.NotEmpty(card)
	if assert.Len(deck.cards, deck.maxSize-1) {
		assert.NotContains(deck.cards, card)
	}
}

func Test_Deck_Place(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("multiple", func(t *testing.T) {
		validIndices := []int{0, 1, 2}
		emptyDeck := Deck{cards: make([]Card, 3), maxSize: 8}
		if assert.NoError(emptyDeck.Place(deck.cards[:3], validIndices)) {
			assert.Equal(deck.cards[:3], emptyDeck.cards[:3])
		}
	})
	t.Run("invalid index", func(t *testing.T) {
		emptyDeck := Deck{cards: make([]Card, 5), maxSize: 10}
		invalidIndices := []int{len(emptyDeck.cards) + 1}
		assert.Error(deck.Place([]Card{}, invalidIndices))
	})
	t.Run("max exceeded", func(t *testing.T) {
		validIndices := []int{0, 1, 2}
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 5}
		partiallyFilledDeck.cards[0] = deck.cards[15]
		assert.Error(deck.Place(deck.cards[:3], validIndices))
	})
	t.Run("mismatched inputs", func(t *testing.T) {
		validIndices := []int{0}
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 3}
		partiallyFilledDeck.cards[0] = deck.cards[15]
		assert.Error(deck.Place(deck.cards[:3], validIndices))
	})
	t.Run("repeated", func(t *testing.T) {
		emptyDeck := Deck{cards: make([]Card, 5), maxSize: 5}
		repeatedIndices := []int{1, 1, 1, 1}
		assert.Error(emptyDeck.Place(deck.cards[:4], repeatedIndices))
	})
}

func TestPlaceTop(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("space available", func(t *testing.T) {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 5}
		if assert.NoError(partiallyFilledDeck.PlaceTop(deck.cards[0])) {
			assert.Equal(deck.cards[0], partiallyFilledDeck.cards[0])
			assert.Len(partiallyFilledDeck.cards, 4)
		}
	})
	t.Run("no space", func(t *testing.T) {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 3}
		assert.Error(partiallyFilledDeck.PlaceTop(deck.cards[0]))
	})
}

func BenchmarkPlaceTop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 5}
		partiallyFilledDeck.PlaceTop(Card{})
	}
}

func TestPlaceBottom(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("space available", func(t *testing.T) {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 5}
		if assert.NoError(partiallyFilledDeck.PlaceBottom(deck.cards[0])) {
			assert.Equal(deck.cards[0], partiallyFilledDeck.cards[len(partiallyFilledDeck.cards)-1])
			assert.Len(partiallyFilledDeck.cards, 4)
		}
	})
	t.Run("no space", func(t *testing.T) {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 3}
		assert.Error(partiallyFilledDeck.PlaceBottom(deck.cards[0]))
	})
}

func TestPlaceRandom(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("space available", func(t *testing.T) {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 5}
		if assert.NoError(partiallyFilledDeck.PlaceRandom(deck.cards[0])) {
			if assert.Len(partiallyFilledDeck.cards, 4) {
				assert.Contains(partiallyFilledDeck.cards, deck.cards[0])
			}
		}
	})
	t.Run("no space", func(t *testing.T) {
		partiallyFilledDeck := Deck{cards: make([]Card, 3), maxSize: 3}
		assert.Error(partiallyFilledDeck.PlaceRandom(deck.cards[0]))
	})
}
