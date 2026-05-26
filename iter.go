//go:build go1.23

package rill

import (
	"iter"
)

// FromSeq converts an iterator into a stream.
// If err is not nil function returns a stream with a single error.
//
// Such function signature allows concise wrapping of functions that return an
// iterator and an error:
//
//	stream := rill.FromSeq(someFunc())
func FromSeq[A any](seq iter.Seq[A], err error) <-chan Try[A] {
	_ = "STUB: not implemented"
	return nil
}

// FromSeq2 converts an iterator of value-error pairs into a stream.
func FromSeq2[A any](seq iter.Seq2[A, error]) <-chan Try[A] { _ = "STUB: not implemented"; return nil }

// ToSeq2 converts an input stream into an iterator of value-error pairs.
//
// This is a blocking ordered function that processes items sequentially.
// It does not return on the first encountered error. Instead, it iterates over all value-error
// pairs, either until the input stream is fully consumed or the loop is broken by the caller.
// So all error handling, if needed, should be done inside the iterator (for-range loop body).
//
// See the package documentation for more information on blocking ordered functions.
func ToSeq2[A any](in <-chan Try[A]) iter.Seq2[A, error] { _ = "STUB: not implemented"; return nil }
