package cards

//Peeker is the interface that wraps the Peek method.
//
//Peek returns a slice of cards at the given indices.
//This does not remove the cards from the underlying cards slice.
type Peeker interface {
	Peek(indices []int) ([]Card, error)
}
