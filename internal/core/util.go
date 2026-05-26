package core

func Drain[A any](in <-chan A) { _ = "STUB: not implemented"; return }

func Discard[A any](in <-chan A) { _ = "STUB: not implemented"; return }

func Buffer[A any](in <-chan A, size int) <-chan A {
	_ = "STUB: not implemented"
	// we use size-1 since 1 additional item is held on the stack (x variable)
	return nil
}
