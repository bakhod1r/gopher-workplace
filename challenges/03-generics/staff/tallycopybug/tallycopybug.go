// Package tallycopybug — Gopher Workplace challenge.
package tallycopybug

// Tally counts occurrences of values for one named bucket.
type Tally[T comparable] struct {
	Name   string
	Counts map[T]int
	Total  int
}

// NewTally returns an initialised Tally.
func NewTally[T comparable](name string) Tally[T] {
	return Tally[T]{Name: name, Counts: make(map[T]int)}
}

// BumpAll records one occurrence of v in every tally in ts.
//
// Each tally's Counts map and its Total must both advance.
//
// Examples:
//
//	BumpAll(ts, "x") twice => every Total is 2
func BumpAll[T comparable](ts []Tally[T], v T) {
	// CHANGE CODE BELOW THIS LINE
	for _, t := range ts {
		t.Counts[v]++
		t.Total++
	}
	// CHANGE CODE ABOVE THIS LINE
}

// Consistent reports whether Total equals the sum of the counts.
func Consistent[T comparable](t Tally[T]) bool {
	sum := 0
	for _, n := range t.Counts {
		sum += n
	}
	return sum == t.Total
}
