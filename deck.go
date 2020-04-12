package main

import (
	"math/rand"
	"time"
)

//BasicDeck is the interface for a Deck of cards.
type BasicDeck interface {
	Shuffler
	Picker
	Dealer
	Placer
	Peeker
}

//Deck implements the BasicDeck interface with additional convenience methods.
type Deck struct {
	cards   []Card
	maxSize int
}

//NewStandardDeck returns a standard Deck of Cards.
//
//If jokers is true the Deck will include jokers.
func NewStandardDeck(jokers bool) Deck {
	var cards []Card
	var maxSize int
	if !jokers {
		cards = make([]Card, 52)
		maxSize = 52
	} else {
		cards = make([]Card, 54)
		maxSize = 54
	}
	allSuits := allSuits()
	allRanks := allRanks()
	for i, suit := range allSuits[:len(allSuits)-2] {
		for j, rank := range allRanks[:len(allRanks)-2] {
			cards[i*13+j] = Card{
				suit: suit,
				rank: rank,
			}
		}
	}
	if jokers {
		cards[52] = Card{suit: allSuits[len(allSuits)-2], rank: allRanks[len(allRanks)-2]}
		cards[53] = Card{suit: allSuits[len(allSuits)-1], rank: allRanks[len(allRanks)-1]}
	}
	return Deck{
		maxSize: maxSize,
		cards:   cards,
	}
}

//Shuffle randomly changes the order of cards in the deck.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	var shiftCount int
	for end := time.Now().Add(time.Millisecond * 100); ; {
		if shiftCount%10 == 0 {
			if time.Now().After(end) {
				break
			}
		}
		index := rand.Intn(len(d.cards))
		d.cards[index], d.cards[len(d.cards)-1] = d.cards[len(d.cards)-1], d.cards[index]
		shiftCount++
	}
}

//Pick returns the cards at the given indices and removes them from the deck.
//
//The indices refer to the state of the deck before any cards are removed.
//Errors if indices are repeated.
func (d *Deck) Pick(indices []int) ([]Card, error) {
	if len(d.cards) == 0 {
		return nil, &OutOfRange{indices: indices}
	}
	if len(indices) == 1 {
		if indices[0] >= len(d.cards) {
			return nil, &OutOfRange{indices: indices}
		}
		card := d.cards[indices[0]]
		d.cards = append(d.cards[:indices[0]], d.cards[indices[0]+1:]...)
		return []Card{card}, nil
	}
	invalid := []int{}
	offsets := make([]int, len(indices))
	indexCount := map[int]int{}
	repeated := []int{}
	for i, a := range indices {
		if a >= len(d.cards) {
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
		picked[i] = d.cards[index-offsets[i]]
		d.cards = append(d.cards[:index-offsets[i]], d.cards[index-offsets[i]+1:]...)
	}
	return picked, nil
}

//Deal returns n hands containing size cards and removes them from the deck.
func (d *Deck) Deal(n, size int) ([]Hand, error) {
	if n*size > len(d.cards) {
		return nil, &NotEnough{requested: n * size, available: len(d.cards)}
	}
	pick := make([]int, n*size)
	for i := 0; i < n*size; i++ {
		pick[i] = i
	}
	cards, err := d.Pick(pick)
	if err != nil {
		return nil, err
	}
	hands := make([]Hand, n)
	for i := range hands {
		hands[i].cards = cards[i*size : (i*size + size)]
		hands[i].maxSize = size
	}
	return hands, nil
}

//PickTop returns the card on top removing it from the deck.
func (d *Deck) PickTop() (Card, error) {
	if len(d.cards) < 1 {
		return Card{}, &NotEnough{requested: 1, available: 0}
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

//PickBottom returns the card at the bottom removing it from the deck.
func (d *Deck) PickBottom() (Card, error) {
	if len(d.cards) < 1 {
		return Card{}, &NotEnough{requested: 1, available: 0}
	}
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card, nil
}

//CardCount returns the number of cards in the deck.
func (d *Deck) CardCount() int {
	return len(d.cards)
}

//HasCard returns true if the deck has a matching card.
func (d *Deck) HasCard(card *Card) bool {
	for _, cardInDeck := range d.cards {
		if cardInDeck.Matches(card) {
			return true
		}
	}
	return false
}

//MaxSize return the maximum number of cards allowed in the deck.
func (d *Deck) MaxSize() int {
	return d.maxSize
}

//Peek returns the cards at the given indices but does not remove them from the deck.
func (d *Deck) Peek(indices []int) ([]Card, error) {
	if len(d.cards) == 0 {
		return nil, &OutOfRange{indices: indices}
	}
	invalid := []int{}
	peeked := make([]Card, len(indices))
	for i, index := range indices {
		if index >= len(d.cards) {
			invalid = append(invalid, index)
		}
		if len(invalid) == 0 {
			peeked[i] = d.cards[index]
		}
	}
	if len(invalid) > 0 {
		return nil, &OutOfRange{indices: invalid}
	}
	return peeked, nil
}

//PeekTop returns the top card without removing it from the deck.
func (d *Deck) PeekTop() (Card, error) {
	if len(d.cards) < 1 {
		return Card{}, &NotEnough{requested: 1, available: 0}
	}
	return d.cards[0], nil
}

//PeekBottom returns the bottom card without removing it from the deck.
func (d *Deck) PeekBottom() (Card, error) {
	if len(d.cards) < 1 {
		return Card{}, &NotEnough{requested: 1, available: 0}
	}
	return d.cards[len(d.cards)-1], nil
}

//PickRandom returns and removes a random card from the deck.
func (d *Deck) PickRandom() (Card, error) {
	if len(d.cards) < 1 {
		return Card{}, &NotEnough{requested: 1, available: 0}
	}
	if len(d.cards) == 1 {
		return d.cards[0], nil
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(d.cards))
	card, err := d.Pick([]int{index})
	if err != nil {
		return Card{}, err
	}
	return card[0], nil
}

//Place inserts cards into the deck at the given indices.
//
//The indices refer to the state of the deck after all cards are inserted.
//
//Place errors if placing cards would exceed the max size,
//an index would be beyond the end of the new deck,
//indices are repeated, or inputs are different sizes.
func (d *Deck) Place(cards []Card, indices []int) error {
	if len(cards) != len(indices) {
		return &MismatchedInputs{inputs: []string{"cards", "indices"}}
	}
	if d.maxSize < len(d.cards)+len(cards) {
		return &NotEnough{requested: len(cards), available: d.maxSize - len(d.cards)}
	}
	repeated := []int{}
	invalid := []int{}
	indexCount := map[int]int{}
	for _, index := range indices {
		if index >= len(d.cards)+len(cards) {
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
	newOrder := make([]Card, len(cards)+len(d.cards))
	for i, card := range cards {
		newOrder[indices[i]] = card
	}
	var insertIndex int
	for _, spot := range newOrder {
		if spot.IsEmpty() {
			spot = d.cards[insertIndex]
			insertIndex++
		}
	}
	d.cards = newOrder
	return nil
}

//PlaceTop places a card at the top of the deck.
func (d *Deck) PlaceTop(card Card) error {
	if len(d.cards) == d.maxSize {
		return &NotEnough{requested: 1, available: 0}
	}
	d.cards = append([]Card{card}, d.cards...)
	return nil
}

//PlaceBottom places a card at the bottom of the deck.
func (d *Deck) PlaceBottom(card Card) error {
	if len(d.cards) == d.maxSize {
		return &NotEnough{requested: 1, available: 0}
	}
	d.cards = append(d.cards, card)
	return nil
}

//PlaceRandom places a card randomly into the deck.
func (d *Deck) PlaceRandom(card Card) error {
	if len(d.cards) == d.maxSize {
		return &NotEnough{requested: 1, available: 0}
	}
	rand.Seed(time.Now().UnixNano())
	spot := rand.Intn(len(d.cards))
	newOrder := make([]Card, len(d.cards)+1)
	newOrder[spot] = card
	copy(newOrder[:spot], d.cards[:spot])
	copy(newOrder[spot+1:], d.cards[spot:])
	d.cards = newOrder
	return nil
}
