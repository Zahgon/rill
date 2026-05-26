package th

import (
	"testing"
)

func FromSlice[A any](slice []A) <-chan A { _ = "STUB: not implemented"; return nil }

func ToSlice[A any](in <-chan A) []A { _ = "STUB: not implemented"; return nil }

func FromRange(start, end int) <-chan int { _ = "STUB: not implemented"; return nil }

func Send[T any](ch chan<- T, items ...T) { _ = "STUB: not implemented"; return }

func Sort[A ordered](s []A) { _ = "STUB: not implemented"; return }

func DoConcurrently(ff ...func()) { _ = "STUB: not implemented"; return }

func DoConcurrentlyN(n int, f func(i int)) { _ = "STUB: not implemented"; return }

// Name generates a test name.
// Works the same way as fmt.Sprint, but adds spaces between all arguments.
func Name(args ...any) string { _ = "STUB: not implemented"; return "" }

func TestBothOrderings(t *testing.T, f func(t *testing.T, ord bool)) {
	_ = "STUB: not implemented"
	return
}
