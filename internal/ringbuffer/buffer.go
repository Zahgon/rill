package ringbuffer

const minCap = 16

type Buffer[T any] struct {
	data         []T
	offset, size int
}

func (b *Buffer[T]) Cap() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer[T]) Len() int {
	_ = "STUB: not implemented"

	// write to end
	return 0
}

func (b *Buffer[T]) Write(v T) { _ = "STUB: not implemented"; return }

// read from start
func (b *Buffer[T]) Read() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (b *Buffer[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (b *Buffer[T]) Discard() bool { _ = "STUB: not implemented"; return false }

// let GC do its work

// change the capacity and defragment the buffer
// panics if newCap is less than buf.size
func (b *Buffer[T]) setCap(newCap int) { _ = "STUB: not implemented"; return }

func (b *Buffer[T]) Grow(n int) { _ = "STUB: not implemented"; return }

// enough

// double the capacity

func (b *Buffer[T]) CanShrink() bool { _ = "STUB: not implemented"; return false }

func (b *Buffer[T]) Shrink() { _ = "STUB: not implemented"; return }

func (b *Buffer[T]) Compact() { _ = "STUB: not implemented"; return }

func (b *Buffer[T]) Reset() { _ = "STUB: not implemented"; return }
