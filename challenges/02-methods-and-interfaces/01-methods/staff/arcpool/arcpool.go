// Package arcpool — Gopher Workplace challenge.
package arcpool

// ARC simulates Adaptive Replacement Cache lists tracking.
type ARC struct {
	T1 map[int]bool // recently used
	T2 map[int]bool // frequently used
}

// Access simulates accessing an item.
// If in T1, it moves to T2. If not in T1 or T2, it goes to T1.
func (a *ARC) Access(key int) {
	// TODO(candidate): implement the logic above.
	panic("not implemented")
}
