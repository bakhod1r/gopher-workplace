// Package hyperlogl — Gopher Workplace challenge.
package hyperlogl

// HLL simulates cardinality estimation by tracking max leading zeros.
type HLL struct {
	maxZeros int
}

func leadingZeros(v uint32) int {
	// mock: just return v % 5 for simulation purposes
	return int(v % 5)
}

// Add updates maxZeros if the new item has more.
func (h *HLL) Add(hash uint32) {
	// TODO(candidate): zeros := leadingZeros(hash); if zeros > h.maxZeros { h.maxZeros = zeros }
	panic("not implemented")
}
