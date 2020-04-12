package main

//Picker is the interface that wraps the Pick method.
//
//Pick removes the cards at the given indicies from the underlying cards slice.
//The cards are returned if all of the indicies are valid otherwise it errors.
type Picker interface {
	Pick(indices []int) ([]Card, error)
}
