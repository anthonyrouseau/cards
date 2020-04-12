package main

import (
	"math/rand"
	"time"
)

//Hand is a set of cards generally held by a player.
type Hand struct {
	cards   []Card
	maxSize int
}

//Peek returns the cards at the given indices but does not remove them from the hand.
func (h *Hand) Peek(indices []int) ([]Card, error) {
	if len(h.cards) == 0 {
		return nil, &OutOfRange{indices: indices}
	}
	invalid := []int{}
	peeked := make([]Card, len(indices))
	for i, index := range indices {
		if index >= len(h.cards) {
			invalid = append(invalid, index)
		}
		if len(invalid) == 0 {
			peeked[i] = h.cards[index]
		}
	}
	if len(invalid) > 0 {
		return nil, &OutOfRange{indices: invalid}
	}
	return peeked, nil
}

//Place inserts cards into the hand at the given indices.
//
//The indices refer to the state of the hand after all cards are inserted.
//
//Place errors if placing cards would exceed the max size,
//an index would be beyond the end of the new hand,
//indices are repeated, or inputs are different sizes.
func (h *Hand) Place(cards []Card, indices []int) error {
	if len(cards) != len(indices) {
		return &MismatchedInputs{inputs: []string{"cards", "indices"}}
	}
	if h.maxSize < len(h.cards)+len(cards) {
		return &NotEnough{requested: len(cards), available: h.maxSize - len(h.cards)}
	}
	repeated := []int{}
	invalid := []int{}
	indexCount := map[int]int{}
	for _, index := range indices {
		if index >= len(h.cards)+len(cards) {
			invalid = append(invalid, index)
		} else {
			indexCount[index]++
			if val, ok := indexCount[index]; ok && val == 2 {
				repeated = append(repeated, index)
			}
		}
	}
	if len(invalid) > 0 {
		return &OutOfRange{indices: invalid}
	}
	if len(repeated) > 0 {
		return &RepeatedIndex{indices: repeated}
	}
	newOrder := make([]Card, len(cards)+len(h.cards))
	for i, card := range cards {
		newOrder[indices[i]] = card
	}
	var insertIndex int
	for _, spot := range newOrder {
		if spot.IsEmpty() {
			spot = h.cards[insertIndex]
			insertIndex++
		}
	}
	h.cards = newOrder
	return nil
}

//Pick returns the cards at the given indices and removes them from the hand.
//
//The indices refer to the state of the hand before any cards are removed.
//Errors if indices are repeated.
func (h *Hand) Pick(indices []int) ([]Card, error) {
	if len(h.cards) == 0 {
		return nil, &OutOfRange{indices: indices}
	}
	if len(indices) == 1 {
		if indices[0] >= len(h.cards) {
			return nil, &OutOfRange{indices: indices}
		}
		card := h.cards[indices[0]]
		h.cards = append(h.cards[:indices[0]], h.cards[indices[0]+1:]...)
		return []Card{card}, nil
	}
	invalid := []int{}
	offsets := make([]int, len(indices))
	indexCount := map[int]int{}
	repeated := []int{}
	for i, a := range indices {
		if a >= len(h.cards) {
			invalid = append(invalid, a)
		} else {
			indexCount[a]++
			if val, ok := indexCount[a]; ok && val == 2 {
				repeated = append(repeated, a)
			}
		}
		for j, b := range indices[i+1:] {
			if a < b {
				offsets[j+i+1]++
			}
		}

	}
	if len(invalid) > 0 {
		return nil, &OutOfRange{indices: invalid}
	}
	if len(repeated) > 0 {
		return nil, &RepeatedIndex{indices: repeated}
	}
	picked := make([]Card, len(indices))
	for i, index := range indices {
		picked[i] = h.cards[index-offsets[i]]
		h.cards = append(h.cards[:index-offsets[i]], h.cards[index-offsets[i]+1:]...)
	}
	return picked, nil
}

//PickRandom returns and removes a random card from the hand.
func (h *Hand) PickRandom() (Card, error) {
	if len(h.cards) < 1 {
		return Card{}, &NotEnough{requested: 1, available: 0}
	}
	if len(h.cards) == 1 {
		return h.cards[0], nil
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(h.cards))
	card, err := h.Pick([]int{index})
	if err != nil {
		return Card{}, err
	}
	return card[0], nil
}

//CardCount returns the number of cards in the hand.
func (h *Hand) CardCount() int {
	return len(h.cards)
}

//MaxSize returns the maximum number of cards allowed in the hand.
func (h *Hand) MaxSize() int {
	return h.maxSize
}

//HasCard returns true if the deck has a matching card.
func (h *Hand) HasCard(card *Card) bool {
	for _, cardInDeck := range h.cards {
		if cardInDeck.Matches(card) {
			return true
		}
	}
	return false
}
