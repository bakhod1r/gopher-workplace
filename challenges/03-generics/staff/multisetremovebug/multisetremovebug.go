// Package multisetremovebug — Gopher Workplace challenge.
package multisetremovebug

// Multiset counts occurrences of comparable values.
type Multiset[T comparable] struct {
	m map[T]int
	n int
}

// Remove takes one occurrence of v out of the multiset.
// It reports whether an occurrence was present.
func (s *Multiset[T]) Remove(v T) bool {
	// CHANGE CODE BELOW THIS LINE
	_, ok := s.m[v]
	if !ok {
		return false
	}
	delete(s.m, v)
	s.n--
	return true
	// CHANGE CODE ABOVE THIS LINE
}

// Add puts one more occurrence of v into the multiset.
func (s *Multiset[T]) Add(v T) {
	if s.m == nil {
		s.m = make(map[T]int)
	}
	s.m[v]++
	s.n++
}

// Count reports how many occurrences of v remain.
func (s *Multiset[T]) Count(v T) int {
	return s.m[v]
}

// Len reports the total number of occurrences.
func (s *Multiset[T]) Len() int {
	return s.n
}

// Distinct reports how many different values are present.
func (s *Multiset[T]) Distinct() int {
	return len(s.m)
}
