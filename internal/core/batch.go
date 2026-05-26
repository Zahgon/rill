package core

import (
	"time"
)

// Batch groups items from an input channel into batches based on a maximum size and a timeout.
// A batch is emitted when it reaches the maximum size, the timeout expires, or the input channel closes.
// This function never emits empty batches. The timeout countdown starts when the first item is added to a new batch.
// To emit batches only when full, set the timeout to -1. Zero timeout is not supported and will panic.
func Batch[A any](in <-chan A, size int, timeout time.Duration) <-chan []A {
	_ = "STUB: not implemented"
	return nil
}

// infinite timeout

// finite timeout

// consume a tick that might have been sent while we were flushing

// timeout

// end of input

// got new item

// we've just started collecting a new batch.
// start the timer to flush the batch after the timeout.

// batch is full

// Unbatch is the inverse of Batch. It takes a channel of batches and emits individual items.
func Unbatch[A any](in <-chan []A) <-chan A { _ = "STUB: not implemented"; return nil }
