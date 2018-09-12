package cards

var (
	Suits = []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	Ranks = []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
)

type Card struct {
	Suit  string
	Rank  string
	Value int
}

type Deck struct {
	Cards []Card
	Size  int
}

func NewDeck() Deck {
	var cards []Card
	for _, suit := range Suits {
		for j, rank := range Ranks {
			cards = append(cards, Card{Suit: suit, Rank: rank, Value: j + 1})
		}
	}
	return Deck{Size: 1, Cards: cards}

}
