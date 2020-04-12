package main

import (
	"log"

	"github.com/anthonyrouseau/cards/cards"
)

func main() {
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
