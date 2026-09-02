// Package bloomifc — Gopher Workplace challenge.
package bloomifc

// Filter reports probable membership.
type Filter interface {
	Add(key string)
	MayContain(key string) bool
}

// Bloom is a two-hash Bloom filter over a fixed bitset.
type Bloom struct {
	bits []byte
	m    uint32 // number of bits
}

// NewBloom returns a filter with m bits (rounded up to a byte boundary).
func NewBloom(m int) *Bloom {
	if m < 8 {
		m = 8
	}
	return &Bloom{bits: make([]byte, (m+7)/8), m: uint32((m + 7) / 8 * 8)}
}

// hash1 is FNV-1a.
func hash1(key string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])
		h *= 16777619
	}
	return h
}

// hash2 is a djb2 variant.
func hash2(key string) uint32 {
	var h uint32 = 5381
	for i := 0; i < len(key); i++ {
		h = h*33 + uint32(key[i])
	}
	return h
}

// Add records a key.
//
// Examples:
//
//	Add("a") => MayContain("a") is true
func (b *Bloom) Add(key string) {
	// TODO(candidate): set the bit for each hash.
	panic("not implemented")
}

// MayContain reports whether the key may have been added. False means it
// definitely was not.
func (b *Bloom) MayContain(key string) bool {
	// TODO(candidate): both bits must be set.
	panic("not implemented")
}

// FilterMissing returns the keys the filter says are definitely absent.
func FilterMissing(f Filter, keys []string) []string {
	// TODO(candidate): keep the keys MayContain rejects.
	panic("not implemented")
}
