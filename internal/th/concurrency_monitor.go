package th

import (
	"sync"
	"time"
)

// ConcurrencyMonitor measures the maximum concurrency level reached by goroutines.
// It enforces maximum possible concurrency by requiring each goroutine to call Inc() at the start and Dec() at the end of its execution.
// Goroutines calling Inc() are blocked until the concurrency level remains stable for a specified time window, ensuring that concurrency peaks are accurately captured.
// The highest level of concurrency observed can be retrieved using the Max() method.
type ConcurrencyMonitor struct {
	cond    *sync.Cond
	current int
	max     int

	target int
	window time.Duration

	lastChangeAt time.Time
	timer        *time.Timer
	timerFired   bool
}

func NewConcurrencyMonitor(window time.Duration) *ConcurrencyMonitor {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConcurrencyMonitor) Inc() { _ = "STUB: not implemented"; return }

// block all goroutines unless "window" has passed since the last counter change

func (c *ConcurrencyMonitor) Dec() { _ = "STUB: not implemented"; return }

func (c *ConcurrencyMonitor) Reset() int { _ = "STUB: not implemented"; return 0 }

func (c *ConcurrencyMonitor) Max() int { _ = "STUB: not implemented"; return 0 }
