// Package versionsnapaliasbug — Gopher Workplace challenge.
package versionsnapaliasbug

// Store is a list with restorable snapshots.
type Store[T any] struct {
	live  []T
	snaps [][]T
}

// Restore replaces the live list with the contents of snapshot id.
// The snapshot itself is left untouched and stays restorable.
func (s *Store[T]) Restore(id int) bool {
	// CHANGE CODE BELOW THIS LINE
	if id < 0 || id >= len(s.snaps) {
		return false
	}
	s.live = s.snaps[id]
	return true
	// CHANGE CODE ABOVE THIS LINE
}

// Snapshot records the current live list and returns its id.
func (s *Store[T]) Snapshot() int {
	c := make([]T, len(s.live))
	copy(c, s.live)
	s.snaps = append(s.snaps, c)
	return len(s.snaps) - 1
}

// Append adds v to the live list.
func (s *Store[T]) Append(v T) {
	s.live = append(s.live, v)
}

// Set overwrites the live element at index i.
func (s *Store[T]) Set(i int, v T) bool {
	if i < 0 || i >= len(s.live) {
		return false
	}
	s.live[i] = v
	return true
}

// Items returns a copy of the live list.
func (s *Store[T]) Items() []T {
	out := make([]T, len(s.live))
	copy(out, s.live)
	return out
}
