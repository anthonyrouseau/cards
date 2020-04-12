package main

import (
	"fmt"
)

//OutOfRange signals an index out of range error.
//
//e.g. picking an index >= the length of the deck.
type OutOfRange struct {
	indices []int
}

func (e *OutOfRange) Error() string {
	return fmt.Sprintf("Indices %v are not in range.", e.indices)
}

//RepeatedIndex signals duplicates of an index
//
//e.g. trying to get the 1 index multiple times []int{1,1,1,1}
type RepeatedIndex struct {
	indices []int
}

func (e *RepeatedIndex) Error() string {
	return fmt.Sprintf("Indices %v were repeated.", e.indices)
}

//NotEnough signals insufficient resources to complete the action.
//
//e.g. trying to deal 20 cards when only 5 are in the deck.
type NotEnough struct {
	requested int
	available int
}

func (e *NotEnough) Error() string {
	return fmt.Sprintf("Requested %d but only %d available.", e.requested, e.available)
}

//MismatchedInputs signals that the values provided are not valid together.
//
//e.g. trying to insert 3 cards into a deck but providing only one index.
type MismatchedInputs struct {
	inputs []string
}

func (e *MismatchedInputs) Error() string {
	return fmt.Sprintf("The combination of inputs %v are not valid.", e.inputs)
}
