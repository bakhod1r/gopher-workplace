// Package mapvalueretain — Gopher Workplace challenge.
package mapvalueretain

// Index stores the n bytes of batch at offset off under key.
//
// The map outlives the batch, so the stored value must own its bytes: a
// view keeps the entire batch reachable for as long as the entry lives.
//
// Examples:
//
//	Index(m, "a", batch, 0, 4) => m["a"] is a 4-byte copy
func Index(m map[string][]byte, key string, batch []byte, off, n int) {
	// CHANGE CODE BELOW THIS LINE
	if m == nil || off < 0 || n < 0 || off+n > len(batch) {
		return
	}
	m[key] = batch[off : off+n : off+n]
	// CHANGE CODE ABOVE THIS LINE
}
