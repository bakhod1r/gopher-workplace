// Package arenastrings — Gopher Workplace challenge.
package arenastrings

import "unsafe"

// Intern returns one string per span of arena, for the caller to keep.
//
// arena is a scratch region the caller refills between batches, so the
// strings must own their bytes. Copying them into one block keeps the cost
// to a single allocation for the whole batch.
//
// Examples:
//
//	Intern([]byte("abcd"), [][2]int{{0,2},{2,4}}) => []string{"ab", "cd"}
func Intern(arena []byte, spans [][2]int) []string {
	// CHANGE CODE BELOW THIS LINE
	if len(spans) == 0 {
		return nil
	}
	out := make([]string, 0, len(spans))
	for _, sp := range spans {
		lo, hi := sp[0], sp[1]
		if lo < 0 || hi > len(arena) || lo > hi {
			out = append(out, "")
			continue
		}
		if lo == hi {
			out = append(out, "")
			continue
		}
		out = append(out, unsafe.String(unsafe.SliceData(arena[lo:hi]), hi-lo))
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
