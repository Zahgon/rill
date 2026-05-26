package core

import (
	"sync"
)

// Loop allows to process items from the input channel concurrently using n goroutines.
// If done channel is not nil, it will be closed after all items are processed.
func Loop[A, B any](in <-chan A, done chan<- B, n int, f func(A)) {
	_ = "STUB: not implemented"
	return
}

type orderedValue[A any] struct {
	Value        A
	CanWrite     chan struct{}
	NextCanWrite chan struct{}
}

var canWritePool sync.Pool

func makeCanWriteChan() chan struct{} { _ = "STUB: not implemented"; return nil }

func releaseCanWriteChan(ch chan struct{}) { _ = "STUB: not implemented"; return }

// OrderedLoop is similar to Loop, but it allows to write results to some channel in the same order as items were read from the input.
// If done channel is not nil, it will be closed after all items are processed.
// Special "canWrite" channel is passed to user's function f. Typical f function looks like this:
// - Do some processing (this part is executed concurrently).
// - Read from canWrite channel exactly once. This step is required. Otherwise, behavior is undefined.
// - Write result of the processing somewhere. This step is optional.
// This way processing is done concurrently, but results are written in order.
func OrderedLoop[A, B any](in <-chan A, done chan<- B, n int, f func(a A, canWrite <-chan struct{})) {
	_ = "STUB: not implemented"
	return
}

// High level idea:
// Each item holds its own canWrite channel and a reference to the next item's canWrite channel.
// After item is processed and written, it sends a signal to the next item that it can also be written.

// first item can be written immediately

// ForEach is a blocking function that processes input channel concurrently using n goroutines
func ForEach[A any](in <-chan A, n int, f func(A)) { _ = "STUB: not implemented"; return }
