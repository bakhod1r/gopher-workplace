// Package versionedgen — Gopher Workplace challenge.
package versionedgen

// Versioned keeps the history of a value of T.
// Its zero value has no history.
type Versioned[T any] struct {
	history []T
}

// Set records a new value, keeping the previous one.
func (v *Versioned[T]) Set(value T) {
	// TODO(candidate): record the new value in the history.
	panic("not implemented")
}

// Get returns the current value and true, or the zero value
// and false when nothing was ever set.
func (v *Versioned[T]) Get() (T, bool) {
	// TODO(candidate): report the latest value.
	panic("not implemented")
}

// Undo drops the most recent value, reporting whether one was
// dropped.
func (v *Versioned[T]) Undo() bool {
	// TODO(candidate): drop the most recent value.
	panic("not implemented")
}

// Versions returns how many values have been recorded.
func (v *Versioned[T]) Versions() int {
	// TODO(candidate): report the history length.
	panic("not implemented")
}
