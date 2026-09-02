// Package compareifc — Gopher Workplace challenge.
package compareifc

// Comparable orders itself against another value.
//
// CompareTo returns <0 when the receiver is smaller, 0 when equal,
// and >0 when the receiver is larger.
type Comparable interface {
	CompareTo(other Comparable) int
}

// Score is a points total.
type Score int

// CompareTo orders two scores.
//
// Examples:
//
//	Score(5).CompareTo(Score(3)) => >0
//	Score(2).CompareTo(Score(2)) => 0
func (s Score) CompareTo(other Comparable) int {
	// TODO(candidate): assert other to Score, then compare.
	panic("not implemented")
}

// Max returns the larger of a and b; a wins a tie.
func Max(a, b Comparable) Comparable {
	// TODO(candidate): use CompareTo.
	panic("not implemented")
}
