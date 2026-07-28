// Package poolreset resets a pooled buffer for reuse. A planted bug clears only
// the length, leaving the data slice non-empty, so a reused buffer keeps stale
// contents.
package poolreset

type Buf struct {
	Data []byte
	Len  int
}

// Reset prepares b for reuse: Len must be 0 AND Data must be emptied (length 0,
// keeping capacity).
func Reset(b *Buf) {
	// CHANGE CODE BELOW THIS LINE
	b.Len = 0
	// CHANGE CODE ABOVE THIS LINE
}
