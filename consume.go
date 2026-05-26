package rill

// ForEach applies a function f to each item in an input stream.
//
// This is a blocking unordered function that processes items concurrently using n goroutines.
// When n = 1, processing becomes sequential, making the function ordered and similar to a regular for-range loop.
//
// See the package documentation for more information on blocking unordered functions and error handling.
func ForEach[A any](in <-chan Try[A], n int, f func(A) error) error {
	_ = "STUB: not implemented"
	return nil
}

// drain

// Err returns the first error encountered in the input stream or nil if there were no errors.
//
// This is a blocking ordered function that processes items sequentially.
// See the package documentation for more information on blocking ordered functions and error handling.
func Err[A any](in <-chan Try[A]) error { _ = "STUB: not implemented"; return nil }

// First returns the first item or error encountered in the input stream, whichever comes first.
// The found return flag is set to false if the stream was empty, otherwise it is set to true.
//
// This is a blocking ordered function that processes items sequentially.
// See the package documentation for more information on blocking ordered functions and error handling.
func First[A any](in <-chan Try[A]) (value A, found bool, err error) {
	_ = "STUB: not implemented"
	return *new(A), false, nil
}

// Any checks if there is an item in the input stream that satisfies the condition f.
// This function returns true as soon as it finds such an item. Otherwise, it returns false.
//
// Any is a blocking unordered function that processes items concurrently using n goroutines.
// When n = 1, processing becomes sequential, making the function ordered.
//
// See the package documentation for more information on blocking unordered functions and error handling.
func Any[A any](in <-chan Try[A], n int, f func(A) (bool, error)) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// drain

// All checks if all items in the input stream satisfy the condition f.
// This function returns false as soon as it finds an item that does not satisfy the condition. Otherwise, it returns true,
// including the case when the stream was empty.
//
// This is a blocking unordered function that processes items concurrently using n goroutines.
// When n = 1, processing becomes sequential, making the function ordered.
//
// See the package documentation for more information on blocking unordered functions and error handling.
func All[A any](in <-chan Try[A], n int, f func(A) (bool, error)) (bool, error) {
	_ = "STUB: not implemented"
	// Idea: x && y && z is the same as !(!x || !y || !z)
	// So we can use Any with a negated condition to implement All
	return false, nil
}

// negate

// negate
