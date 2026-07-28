// Package fnv1a computes the 32-bit FNV-1a hash.
// A planted bug swaps the xor/multiply order (that is FNV-1, not FNV-1a).
package fnv1a

const (
	offset32 = 2166136261
	prime32  = 16777619
)

// Hash returns the FNV-1a 32-bit hash of data: for each byte, XOR then multiply.
func Hash(data []byte) uint32 {
	h := uint32(offset32)
	for _, b := range data {
		// CHANGE CODE BELOW THIS LINE
		h = (h * prime32) ^ uint32(b)
		// CHANGE CODE ABOVE THIS LINE
	}
	return h
}
