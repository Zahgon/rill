package rill

// Merge performs a fan-in operation on the list of input channels, returning a single output channel.
// The resulting channel will contain all items from all inputs,
// and will be closed when all inputs are fully consumed.
//
// This is a non-blocking function that processes items from each input sequentially.
//
// See the package documentation for more information on non-blocking functions and error handling.
func Merge[A any](ins ...<-chan A) <-chan A { _ = "STUB: not implemented"; return nil }

// Split2 divides the input stream into two output streams based on the predicate function f:
// The splitting behavior is determined by the boolean return value of f. When f returns true, the item is sent to the outTrue stream,
// otherwise it is sent to the outFalse stream. In case of any error, the item is sent to both output streams.
// Both output streams must be consumed independently to avoid deadlocks.
//
// This is a non-blocking unordered function that processes items concurrently using n goroutines.
// An ordered version of this function, [OrderedSplit2], is also available.
//
// See the package documentation for more information on non-blocking unordered functions and error handling.
func Split2[A any](in <-chan Try[A], n int, f func(A) (bool, error)) (outTrue <-chan Try[A], outFalse <-chan Try[A]) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OrderedSplit2 is the ordered version of [Split2].
func OrderedSplit2[A any](in <-chan Try[A], n int, f func(A) (bool, error)) (outTrue <-chan Try[A], outFalse <-chan Try[A]) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tee returns two streams that are identical to the input stream (both errors and values).
// Both output streams must be consumed independently to avoid deadlocks.
//
// This is a non-blocking function that processes items in a single goroutine.
// See the package documentation for more information on non-blocking functions and error handling.
//
// If deep copying of values is needed, use [Map] on one or both outputs:
//
//	out1, out2 := rill.Tee(in)
//	out2 = rill.Map(out2, 1, func(x A) (A, error) {
//		return deepCopy(x), nil
//	})
func Tee[A any](in <-chan Try[A]) (<-chan Try[A], <-chan Try[A]) {
	_ = "STUB: not implemented"
	return nil, nil
}
