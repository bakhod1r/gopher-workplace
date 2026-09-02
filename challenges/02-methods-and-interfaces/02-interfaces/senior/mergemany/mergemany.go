// Package mergemany — Gopher Workplace challenge.
package mergemany

// Feed yields ascending values.
type Feed interface {
	Next() (int, bool)
}

// SortedFeed serves an ascending slice.
type SortedFeed struct {
	Data []int
	pos  int
}

// Next yields the next value.
func (f *SortedFeed) Next() (int, bool) {
	// TODO(candidate): yield Data[pos], then advance.
	panic("not implemented")
}

// MergeAll merges any number of ascending feeds into one ascending slice.
//
// Equal values are taken from the earliest feed.
//
// Examples:
//
//	MergeAll([1 4], [2 5], [3]) => [1 2 3 4 5]
func MergeAll(feeds ...Feed) []int {
	// TODO(candidate): hold one head per feed; take the smallest each step.
	panic("not implemented")
}
