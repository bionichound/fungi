package fungi

// Either encodes any computation that could've errored out. It consists of either a
// `Left` or a `Right` where a left represents an error and a right represents a
// successful computation.
//
// This can be used alongisde the `Result` function in order to handle multiple
// functions that can error without as many `if err != nil`.
type Either[A any] struct {
	Left  error
	Right A
}

// Do applies a function that does not fail onto an Either. This means that if the
// either is a successful value, it will pass it in to the provided function. Otherwise,
// the either remains untouched.
func Do[A, B any](m *Either[A], f func(in A) B) *Either[B] {
	var nilB B
	if m.Left != nil {
		return &Either[B]{Left: m.Left, Right: nilB}
	}
	return &Either[B]{Left: nil, Right: f(m.Right)}
}

// Bind applies a function that can fail onto an Either. Similar to the Do function,
// however, it wraps the function return in an either, where the function returns
// a value and an error.
func Bind[A, B any](m *Either[A], f func(in A) (B, error)) *Either[B] {
	var nilB B
	if m.Left != nil {
		return &Either[B]{Left: m.Left, Right: nilB}
	}
	return Result(f(m.Right))
}

// Result takes in the result of any function that follows idiomatic go for error
// handling such that it returns a value, err pair. It then wraps it in an Either
// that can be used to chain multiple functions that may error reducing the amount
// of logic required to handle errors.
func Result[A any](out A, err error) *Either[A] {
	return &Either[A]{Left: err, Right: out}
}
