package core

func FilterMap[A, B any](in <-chan A, n int, f func(A) (B, bool)) <-chan B {
	_ = "STUB: not implemented"
	return nil
}

func OrderedFilterMap[A, B any](in <-chan A, n int, f func(A) (B, bool)) <-chan B {
	_ = "STUB: not implemented"
	return nil
}
