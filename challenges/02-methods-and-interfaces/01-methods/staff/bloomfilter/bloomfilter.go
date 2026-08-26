// Package bloomfilter — Gopher Workplace challenge.
package bloomfilter

// Filter simulates a bloom filter with a simple bitset.
type Filter struct {
	bits [256]bool
}

func hash1(s string) byte { return s[0] }
func hash2(s string) byte { return s[len(s)-1] }

// Add sets the bits for both hashes.
func (f *Filter) Add(item string) {
	if len(item) == 0 {
		return
	}
	// TODO(candidate): f.bits[hash1(item)] = true; f.bits[hash2(item)] = true
	panic("not implemented")
}

// MightContain returns true if both bits are set.
func (f *Filter) MightContain(item string) bool {
	if len(item) == 0 {
		return false
	}
	// TODO(candidate): return f.bits[hash1(item)] && f.bits[hash2(item)]
	panic("not implemented")
}
