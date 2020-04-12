package main

//Dealer is the interface that wraps the Deal method.
//
//Deal returns a slice of n Hands of a given size by selecting cards from the underlying cards slice.
//This removes the vards from the slice.
type Dealer interface {
	Deal(n, size int) ([]*Hand, error)
}
