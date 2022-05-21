// Fungi is a functional programming library for golang.
//
// It provides functions for operations on lists, the creation and manipulation of lazy
// (iterable) collections and the handling of errors and possibly missing fields through
// the use of Haskell inspired monads.
package fungi

// Map is a function that maps a function over a generic collection, returning a new
// generic collection.
func Map[T, U interface{}](ls []T, fun func(i T) U) []U {
	size := len(ls)
	result := make([]U, size)

	for n, i := range ls {
		result[n] = fun(i)
	}
	return result
}

// Fold is a function that 'reduces' a generic collection into a new generic value using
// a reduction/fold function and an initial value.
func Fold[A, B interface{}](r func(item A, acc B) B, init B, ls []A) B {
	acc := init

	for _, item := range ls {
		acc = r(item, acc)
	}

	return acc
}

// Filter returns a subset of an original generic collection using a predicate function.
func Filter[A interface{}](pred func(item A) bool, ls []A) (result []A) {
	for _, item := range ls {
		if pred(item) {
			result = append(result, item)
		}
	}

	return result
}

// Includes is a function that returns whether a particular value is contained within a
// generic collection.
func Includes[A comparable](ls []A, item A) (ret bool) {
	for _, i := range ls {
		if i == item {
			return true
		}
	}
	return ret
}
