package cards

//Shuffler is the interface that wraps the Shuffle method.
//
//Shuffle randomly sets the order of the underlying cards slice.
type Shuffler interface {
	Shuffle()
}
