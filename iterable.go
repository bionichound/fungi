package fungi

// Iterable is an interface that represents a lazy collection that can be resolved on a
// per-item basis. In order to implement the iterable interface, a single method needs to
// be defined in which the next item in a series is returned.
type Iterable[A any] interface {
	Next() A
}

// Numbers is an iterable that returns a sequence of integer numbers. It has one field
// which specifies the first and current number of the sequence.
type Numbers struct {
	Current int
}

func (n *Numbers) Next() (val int) {
	val = n.Current
	n.Current++
	return
}

// Take is a utility function that creates lists from an iterable.
//
// Take requires a source iterable and a number of items to take from it.
// Returns a list of size count taken from the supplied iterable.
func Take[A Iterable[B], B any](src A, count int) []B {
	result := make([]B, count)
	for i := 0; i < count; i++ {
		result[i] = src.Next()
	}
	return result
}
