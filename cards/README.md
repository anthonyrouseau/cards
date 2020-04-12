# Cards

This package provides structs and methods useful for building card games.

## Installation

To install this package use the command:

  `go get github.com/anthonyrouseau/games/cards`

## Example 

````Go
package main

import (
	"log"

	"github.com/anthonyrouseau/games/cards"
)

func main() {
    //Create a basic 52 card deck without jokers
    deck := cards.NewStandardDeck(false)
    //Randomly shuffles the deck
    deck.Shuffle()
    //Create 3 hands with 5 cards each
	hands, err := deck.Deal(3, 5)
	if err != nil {
		panic(err)
    }
    //Go through each hand and print each card in the hand
	for _, hand := range hands {
		for _, card := range hand.cards {
			log.Println(card)
		}
	}

}
````