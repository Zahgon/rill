package rill

// Try is a container holding a value of type A or an error
type Try[A any] struct {
	Value A
	Error error
}

// Wrap converts a value and/or error into a [Try] container.
// It's a convenience function to avoid creating a [Try] container manually and benefit from type inference.
//
// Such function signature also allows concise wrapping of functions that return a value and an error:
//
//	item := rill.Wrap(strconv.ParseInt("42"))
func Wrap[A any](value A, err error) Try[A] { _ = "STUB: not implemented"; return nil }

// FromSlice converts a slice into a stream.
// If err is not nil function returns a stream with a single error.
//
// Such function signature allows concise wrapping of functions that return a slice and an error:
//
//	stream := rill.FromSlice(someFunc())
func FromSlice[A any](slice []A, err error) <-chan Try[A] { _ = "STUB: not implemented"; return nil }

// ToSlice converts an input stream into a slice.
//
// This is a blocking ordered function that processes items sequentially.
// See the package documentation for more information on blocking ordered functions and error handling.
func ToSlice[A any](in <-chan Try[A]) ([]A, error) { _ = "STUB: not implemented"; return nil, nil }

// FromChan converts a regular channel into a stream.
// Additionally, this function can take an error, that will be added to the output stream alongside the values.
// Either argument can be nil, in which case it is ignored. If both arguments are nil, the function returns nil.
//
// Such function signature allows concise wrapping of functions that return a channel and an error:
//
//	stream := rill.FromChan(someFunc())
func FromChan[A any](values <-chan A, err error) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// error goes first

// FromChans converts a regular channel into a stream.
// Additionally, this function can take a channel of errors, which will be added to
// the output stream alongside the values.
// Either argument can be nil, in which case it is ignored. If both arguments are nil, the function returns nil.
//
// Such function signature allows concise wrapping of functions that return two channels:
//
//	stream := rill.FromChans(someFunc())
func FromChans[A any](values <-chan A, errs <-chan error) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// ToChans splits an input stream into two channels: one for values and one for errors.
// It's an inverse of [FromChans]. Returns two nil channels if the input is nil.
func ToChans[A any](in <-chan Try[A]) (<-chan A, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate is a shorthand for creating streams.
// It provides a more ergonomic way of sending both values and errors to a stream, manages goroutine and channel lifecycle.
//
//	stream := rill.Generate(func(send func(int), sendErr func(error)) {
//		for i := 0; i < 100; i++ {
//			send(i)
//		}
//		sendErr(someError)
//	})
//
// Here's how the same code would look without Generate:
//
//	stream := make(chan rill.Try[int])
//	go func() {
//		defer close(stream)
//		for i := 0; i < 100; i++ {
//			stream <- rill.Try[int]{Value: i}
//		}
//		stream <- rill.Try[int]{Error: someError}
//	}()
func Generate[A any](f func(send func(A), sendErr func(error))) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}
