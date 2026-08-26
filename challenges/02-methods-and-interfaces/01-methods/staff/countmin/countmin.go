// Package countmin — Gopher Workplace challenge.
package countmin

// Sketch simulates a CMS with two rows.
type Sketch struct {
	row1 [256]int
	row2 [256]int
}

func h1(s string) byte { return s[0] }
func h2(s string) byte { return s[len(s)-1] }

// Add increments counts.
func (s *Sketch) Add(item string) {
	if len(item) == 0 {
		return
	}
	// TODO(candidate): s.row1[h1]++, s.row2[h2]++
	panic("not implemented")
}

// Count returns the minimum of the counts.
func (s *Sketch) Count(item string) int {
	if len(item) == 0 {
		return 0
	}
	// TODO(candidate): min(s.row1[h1], s.row2[h2])
	panic("not implemented")
}
