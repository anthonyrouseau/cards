package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Card_Matches(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("matches", func(t *testing.T) {
		assert.True(deck.cards[0].Matches(&deck.cards[0]))
	})
	t.Run("doesn't match", func(t *testing.T) {
		assert.False(deck.cards[0].Matches(&deck.cards[1]))
	})
}

func Test_Card_MatchesRank(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("matches", func(t *testing.T) {
		assert.True(deck.cards[0].MatchesRank(&deck.cards[0]))
	})
	t.Run("doesn't match", func(t *testing.T) {
		assert.False(deck.cards[0].MatchesRank(&deck.cards[1]))
	})
}

func Test_Card_MatchesSuit(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("matches", func(t *testing.T) {
		assert.True(deck.cards[0].MatchesSuit(&deck.cards[0]))
	})
	t.Run("doesn't match", func(t *testing.T) {
		assert.False(deck.cards[0].MatchesSuit(&deck.cards[20]))
	})
}

func Test_Card_MatchesColor(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("matches", func(t *testing.T) {
		assert.True(deck.cards[0].MatchesColor(&deck.cards[0]))
	})
	t.Run("doesn't match", func(t *testing.T) {
		card := &Card{suit: Suit{color: Black}}
		assert.False(card.MatchesColor(&Card{suit: Suit{color: Red}}))
	})
}

func Test_Card_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	deck := NewStandardDeck(false)
	t.Run("empty", func(t *testing.T) {
		card := Card{}
		assert.True(card.IsEmpty())
	})
	t.Run("not empty", func(t *testing.T) {
		assert.False(deck.cards[0].IsEmpty())
	})
}
