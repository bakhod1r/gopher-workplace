// Package popretain — Gopher Workplace challenge.
package popretain

// Job is one queued unit of work.
type Job struct {
	ID  int
	Pad [1024]byte
}

// Pop removes and returns the last element of s.
//
// The shortened slice must not keep the popped job reachable: the element
// is still in the backing array until it is cleared.
//
// Examples:
//
//	Pop([]*Job{a, b}) => b, a slice holding only a
func Pop(s []*Job) (*Job, []*Job) {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		return nil, s
	}
	last := s[len(s)-1]
	return last, s[:len(s)-1]
	// CHANGE CODE ABOVE THIS LINE
}
