package rill

// Map takes a stream of items of type A and transforms them into items of type B using a function f.
// Returns a new stream of transformed items.
//
// This is a non-blocking unordered function that processes items concurrently using n goroutines.
// An ordered version of this function, [OrderedMap], is also available.
//
// See the package documentation for more information on non-blocking unordered functions and error handling.
func Map[A, B any](in <-chan Try[A], n int, f func(A) (B, error)) <-chan Try[B] {
	_ = "STUB: not implemented"
	return nil
}

// OrderedMap is the ordered version of [Map].
func OrderedMap[A, B any](in <-chan Try[A], n int, f func(A) (B, error)) <-chan Try[B] {
	_ = "STUB: not implemented"
	return nil
}

// Filter takes a stream of items of type A and filters them using a predicate function f.
// Returns a new stream of items that passed the filter.
//
// This is a non-blocking unordered function that processes items concurrently using n goroutines.
// An ordered version of this function, [OrderedFilter], is also available.
//
// See the package documentation for more information on non-blocking unordered functions and error handling.
func Filter[A any](in <-chan Try[A], n int, f func(A) (bool, error)) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// never filter out errors

// never filter out errors

// OrderedFilter is the ordered version of [Filter].
func OrderedFilter[A any](in <-chan Try[A], n int, f func(A) (bool, error)) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// never filter out errors

// never filter out errors

// FilterMap takes a stream of items of type A, applies a function f that can filter and transform them into items of type B.
// Returns a new stream of transformed items that passed the filter. This operation is equivalent to a
// [Filter] followed by a [Map].
//
// This is a non-blocking unordered function that processes items concurrently using n goroutines.
// An ordered version of this function, [OrderedFilterMap], is also available.
//
// See the package documentation for more information on non-blocking unordered functions and error handling.
func FilterMap[A, B any](in <-chan Try[A], n int, f func(A) (B, bool, error)) <-chan Try[B] {
	_ = "STUB: not implemented"
	return nil
}

// OrderedFilterMap is the ordered version of [FilterMap].
func OrderedFilterMap[A, B any](in <-chan Try[A], n int, f func(A) (B, bool, error)) <-chan Try[B] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMap takes a stream of items of type A and transforms each item into a new sub-stream of items of type B using a function f.
// Those sub-streams are then flattened into a single output stream, which is returned.
//
// This is a non-blocking unordered function that processes items concurrently using n goroutines.
// An ordered version of this function, [OrderedFlatMap], is also available.
//
// See the package documentation for more information on non-blocking unordered functions and error handling.
func FlatMap[A, B any](in <-chan Try[A], n int, f func(A) <-chan Try[B]) <-chan Try[B] {
	_ = "STUB: not implemented"
	return nil
}

// OrderedFlatMap is the ordered version of [FlatMap].
func OrderedFlatMap[A, B any](in <-chan Try[A], n int, f func(A) <-chan Try[B]) <-chan Try[B] {
	_ = "STUB: not implemented"
	return nil
}

// Catch allows handling errors in the middle of a stream processing pipeline.
// Every error encountered in the input stream is passed to the function f for handling.
//
// The outcome depends on the return value of f:
//   - If f returns nil, the error is considered handled and filtered out from the output stream.
//   - If f returns a non-nil error, the original error is replaced with the result of f.
//
// This is a non-blocking unordered function that handles errors concurrently using n goroutines.
// An ordered version of this function, [OrderedCatch], is also available.
//
// See the package documentation for more information on non-blocking unordered functions and error handling.
func Catch[A any](in <-chan Try[A], n int, f func(error) error) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// error handled, filter out

// error replaced by f(a.Error)

// OrderedCatch is the ordered version of [Catch].
func OrderedCatch[A any](in <-chan Try[A], n int, f func(error) error) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// error handled, filter out

// error replaced by f(a.Error)
