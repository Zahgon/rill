package core

func fastMerge[A any](ins []<-chan A) <-chan A {
	_ = "STUB: not implemented"
	// len(ins) must be between 2 and 5
	return nil
}

func slowMerge[A any](ins []<-chan A) <-chan A { _ = "STUB: not implemented"; return nil }

func Merge[A any](ins ...<-chan A) <-chan A { _ = "STUB: not implemented"; return nil }
