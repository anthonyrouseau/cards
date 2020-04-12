package examples

import (
	"log"

	"github.com/anthonyrouseau/games/cards"
)

//CardsExample runs a function with basic usage of the cards package.
func CardsExample() {
	deck := cards.NewStandardDeck(false)
	deck.Shuffle()
	hands, err := deck.Deal(3, 5)
	if err != nil {
		panic(err)
	}
	for _, hand := range hands {
		for _, card := range hand.cards {
			log.Println(card)
		}
	}

}
