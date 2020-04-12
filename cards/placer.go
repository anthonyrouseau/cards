package cards

//Placer is the interface that wraps the Place method.
//
//Place adds cards at the given indicies to the underlying cards slice.
//Place errors if placing the cards into the slice would not be valid.
type Placer interface {
	Place(cards []Card, indices []int) error
}
