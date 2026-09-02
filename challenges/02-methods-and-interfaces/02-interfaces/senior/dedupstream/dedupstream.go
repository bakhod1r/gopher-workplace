// Package dedupstream — Gopher Workplace challenge.
package dedupstream

// Source yields event ids until drained.
type Source interface {
	Next() (string, bool)
}

// SeenSet reports whether an id was already seen, recording it either way.
type SeenSet interface {
	Seen(id string) bool
}

// SliceSource streams a slice of ids.
type SliceSource struct {
	IDs []string
	pos int
}

// Next yields the next id.
func (s *SliceSource) Next() (string, bool) {
	if s.pos >= len(s.IDs) {
		return "", false
	}
	id := s.IDs[s.pos]
	s.pos++
	return id, true
}

// ExactSet remembers every id it has ever seen.
type ExactSet struct {
	seen map[string]bool
}

// NewExactSet returns an empty exact set.
func NewExactSet() *ExactSet {
	return &ExactSet{seen: make(map[string]bool)}
}

// Seen reports whether id was already recorded, then records it.
func (e *ExactSet) Seen(id string) bool {
	// TODO(candidate): check, then record.
	panic("not implemented")
}

// WindowSet remembers only the last N distinct ids.
type WindowSet struct {
	N     int
	seen  map[string]bool
	order []string
}

// NewWindowSet returns a sliding-window set of size n.
func NewWindowSet(n int) *WindowSet {
	return &WindowSet{N: n, seen: make(map[string]bool, n), order: make([]string, 0, n)}
}

// Seen reports whether id is in the window, then records it, evicting the
// oldest id when the window is full.
func (w *WindowSet) Seen(id string) bool {
	// TODO(candidate): windowed membership with eviction.
	panic("not implemented")
}

// Dedup streams src through set and returns the number of unique events.
//
// Examples:
//
//	Dedup over [a a b] with an exact set => 2
func Dedup(src Source, set SeenSet) int {
	// TODO(candidate): count the ids the set had not seen.
	panic("not implemented")
}
