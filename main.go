package main

import "log"

func main() {
	deck := NewStandardDeck(false)
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
