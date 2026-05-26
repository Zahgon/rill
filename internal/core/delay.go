package core

import (
	"time"
)

func infiniteBuffer[A any](in <-chan A) <-chan A { _ = "STUB: not implemented"; return nil }

type delayedValue[A any] struct {
	Value  A
	SendAt time.Time
}

// Delay postpones the delivery of items from an input channel by a specified duration, maintaining the order.
// Useful for adding delays in processing or simulating latency.
func Delay[A any](in <-chan A, delay time.Duration) <-chan A { _ = "STUB: not implemented"; return nil }

// buffering is needed to freely use sleeps in the loop below
