// Package th provides basic test helpers.
package th

import (
	"testing"
	"time"
)

func ExpectValue[A comparable](t *testing.T, actual A, expected A) {
	_ = "STUB: not implemented"
	return
}

func ExpectValueLTE[A number](t *testing.T, actual A, expected A) {
	_ = "STUB: not implemented"
	return
}

func ExpectValueGTE[A number](t *testing.T, actual A, expected A) {
	_ = "STUB: not implemented"
	return
}

func ExpectValueInDelta[A number](t *testing.T, actual A, expected A, delta A) {
	_ = "STUB: not implemented"
	return
}

func ExpectSlice[A comparable](t *testing.T, actual []A, expected []A) {
	_ = "STUB: not implemented"
	return
}

func ExpectMap[K, V comparable](t *testing.T, actual map[K]V, expected map[K]V) {
	_ = "STUB: not implemented"
	return
}

type number interface {
	~int | ~int64
}

type ordered interface {
	~int | ~int64 | ~string
}

func ExpectSorted[T ordered](t *testing.T, arr []T) { _ = "STUB: not implemented"; return }

func ExpectUnsorted[T ordered](t *testing.T, arr []T) { _ = "STUB: not implemented"; return }

func ExpectDrainedChan[A any](t *testing.T, ch <-chan A) { _ = "STUB: not implemented"; return }

func ExpectNeverClosedChan[A any](t *testing.T, ch <-chan A, waitFor time.Duration) {
	_ = "STUB: not implemented"
	return
}

func ExpectHang(t *testing.T, waitFor time.Duration, f func()) { _ = "STUB: not implemented"; return }

func ExpectNotHang(t *testing.T, waitFor time.Duration, f func()) {
	_ = "STUB: not implemented"
	return
}

func ExpectError(t *testing.T, err error, message string) { _ = "STUB: not implemented"; return }

func ExpectNoError(t *testing.T, err error) { _ = "STUB: not implemented"; return }

func ExpectNotPanic(t *testing.T, f func()) { _ = "STUB: not implemented"; return }
