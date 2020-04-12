package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hand_Pick(t *testing.T) {
	assert := assert.New(t)
	t.Run("multiple", func(t *testing.T) {
		deck := NewStandardDeck(false)
		hands, err := deck.Deal(1, 5)
		if assert.NoError(err) {
			hand := hands[0]
			validIndices := []int{0, 1, 2}
			expected := make([]Card, len(validIndices))
			for i, j := range validIndices {
				expected[i] = hand.cards[j]
			}
			got, err := hand.Pick(validIndices)
			if assert.NoError(err) {
				for _, card := range got {
					assert.Contains(expected, card)
				}
				assert.Len(hand.cards, hand.maxSize-len(validIndices))
				for _, handCard := range hand.cards {
					for _, gotCard := range got {
						assert.NotEqual(handCard, gotCard)
					}
				}

			}
		}
	})
	t.Run("single", func(t *testing.T) {
		deck := NewStandardDeck(false)
		hands, err := deck.Deal(1, 5)
		assert.NoError(err)
		hand := hands[0]
		validIndices := []int{1}
		expected := hand.cards[1]
		got, err := hand.Pick(validIndices)
		if assert.NoError(err) {
			assert.Equal(expected, got[0])
		}
		if assert.Len(hand.cards, hand.maxSize-len(validIndices)) {
			for _, handCard := range hand.cards {
				for _, gotCard := range got {
					assert.NotEqual(handCard, gotCard)
				}
			}
		}
	})
	t.Run("invalid", func(t *testing.T) {
		deck := NewStandardDeck(false)
		hands, err := deck.Deal(1, 5)
		assert.NoError(err)
		hand := hands[0]
		invalidIndices := []int{len(hand.cards)}
		_, err = hand.Pick(invalidIndices)
		assert.Error(err)
	})
	t.Run("repeated", func(t *testing.T) {
		deck := NewStandardDeck(false)
		hands, err := deck.Deal(1, 5)
		assert.NoError(err)
		hand := hands[0]
		repeatedIndices := []int{1, 1, 1, 1}
		_, err = hand.Pick(repeatedIndices)
		assert.Error(err)
	})
}

func Test_Hand_CardCount(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	hands, err := deck.Deal(1, 5)
	if assert.NoError(err) {
		hand := hands[0]
		assert.Len(hand.cards, hand.CardCount())
	}
}

func Test_Hand_HasCard(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	hands, err := deck.Deal(1, 5)
	if assert.NoError(err) {
		hand := hands[0]
		t.Run("in hand", func(t *testing.T) {
			card := hand.cards[1]
			assert.True(hand.HasCard(&card))
		})
		t.Run("not in hand", func(t *testing.T) {
			card := &Card{}
			assert.False(hand.HasCard(card))
		})
	}
}

func Test_Hand_MaxSize(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	hands, err := deck.Deal(1, 5)
	if assert.NoError(err) {
		hand := hands[0]
		assert.Equal(hand.maxSize, hand.MaxSize())
	}
}

func Test_Hand_Peek(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	hands, err := deck.Deal(1, 5)
	if assert.NoError(err) {
		hand := hands[0]
		t.Run("multiple", func(t *testing.T) {
			validIndices := []int{0, 1, 2}
			expected := hand.cards[:3]
			got, err := hand.Peek(validIndices)
			if assert.NoError(err) {
				for _, card := range got {
					assert.Contains(expected, card)
				}
			}
			assert.Len(hand.cards, hand.maxSize)
		})
		t.Run("single", func(t *testing.T) {
			validIndices := []int{1}
			expected := hand.cards[1]
			got, err := hand.Peek(validIndices)
			if assert.NoError(err) {
				assert.Equal(expected, got[0])
			}
			assert.Len(hand.cards, hand.maxSize)
		})
		t.Run("invalid", func(t *testing.T) {
			invalidIndices := []int{len(hand.cards)}
			_, err := hand.Peek(invalidIndices)
			assert.Error(err)
		})
	}
}

func Test_Hand_PickRandom(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	hands, err := deck.Deal(1, 5)
	if assert.NoError(err) {
		hand := hands[0]
		card, err := hand.PickRandom()
		if assert.NoError(err) {
			assert.NotEmpty(card)
			assert.Len(hand.cards, hand.maxSize-1)
			assert.NotContains(hand.cards, card)
		}
	}
}

func Test_Hand_Place(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("multiple", func(t *testing.T) {
		hands, err := deck.Deal(1, 1)
		if assert.NoError(err) {
			hand := hands[0]
			hand.maxSize = 5
			validIndices := []int{0, 1, 2}
			if assert.NoError(hand.Place(deck.cards[:3], validIndices)) {
				assert.Equal(deck.cards[:3], hand.cards[:3])
			}
		}
	})
	t.Run("invalid index", func(t *testing.T) {
		hands, err := deck.Deal(1, 1)
		if assert.NoError(err) {
			hand := hands[0]
			hand.maxSize = 5
			invalidIndices := []int{len(hand.cards) + 1}
			assert.Error(hand.Place([]Card{}, invalidIndices))
		}
	})
	t.Run("max exceeded", func(t *testing.T) {
		hands, err := deck.Deal(1, 1)
		if assert.NoError(err) {
			hand := hands[0]
			validIndices := []int{0}
			assert.Error(hand.Place(deck.cards[:1], validIndices))
		}
	})
	t.Run("mismatched inputs", func(t *testing.T) {
		hands, err := deck.Deal(1, 1)
		if assert.NoError(err) {
			hand := hands[0]
			hand.maxSize = 5
			validIndices := []int{0}
			assert.Error(hand.Place(deck.cards[:3], validIndices))
		}
	})
	t.Run("repeated", func(t *testing.T) {
		hands, err := deck.Deal(1, 1)
		if assert.NoError(err) {
			hand := hands[0]
			hand.maxSize = 5
			repeatedIndices := []int{1, 1, 1, 1}
			assert.Error(hand.Place(deck.cards[:4], repeatedIndices))
		}
	})
}
