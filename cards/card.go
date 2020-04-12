package cards

//Card is a standard playing card
type Card struct {
	suit Suit
	rank Rank
}

//Matches returns true if the cards match rank and suit.
func (c *Card) Matches(other *Card) bool {
	return c.MatchesRank(other) && c.MatchesSuit(other)
}

//MatchesRank returns true if the cards match rank.
func (c *Card) MatchesRank(other *Card) bool {
	return c.rank == other.rank
}

//MatchesSuit returns true if the cards match suit.
func (c *Card) MatchesSuit(other *Card) bool {
	return c.suit == other.suit
}

//MatchesColor returns true if the cards match color.
func (c *Card) MatchesColor(other *Card) bool {
	return c.suit.color == other.suit.color
}

//IsEmpty returns true if none of the cards fields are set.
func (c *Card) IsEmpty() bool {
	return c.rank.isEmpty() && c.suit.isEmpty()
}

//Rank of a card, 13 in standard 52 card deck.
type Rank int

func (r *Rank) isEmpty() bool {
	return int(*r) == 0
}

//Rank Values
const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	LittleJoker
	BigJoker
)

func allRanks() []Rank {
	return []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, LittleJoker, BigJoker}
}

//Suit of a card, i.e. Clubs, Diamonds, Hearts, or Spades and its color, red or black
type Suit struct {
	color SuitColor
	name  SuitName
}

func (s *Suit) isEmpty() bool {
	return string(s.color) == "" && string(s.name) == ""
}

func allSuits() []Suit {
	return []Suit{{color: Black, name: Clubs}, {color: Black, name: Spades}, {color: Red, name: Diamonds}, {color: Red, name: Hearts}, {color: Black, name: Joker}, {color: Red, name: Joker}}
}

//SuitName is the suit name e.g. Clubs
type SuitName string

//SuitName values
const (
	Joker    SuitName = "joker"
	Clubs    SuitName = "clubs"
	Diamonds SuitName = "diamonds"
	Hearts   SuitName = "hearts"
	Spades   SuitName = "spades"
)

//SuitColor is the color of the suit e.g. Red
type SuitColor string

//SuitColor values
const (
	Red   SuitColor = "red"
	Black SuitColor = "black"
)
