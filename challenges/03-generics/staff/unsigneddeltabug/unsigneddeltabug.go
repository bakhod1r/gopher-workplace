// Package unsigneddeltabug — Gopher Workplace challenge.
package unsigneddeltabug

// Unsigned is the set of unsigned integer types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Deltas returns the per-sample increase of a monotonic counter.
// A counter reset — a sample lower than its predecessor — contributes the new sample itself.
//
// Examples:
//
//	Deltas([]uint8{10, 250, 5}) => []uint8{10, 240, 5}
func Deltas[T Unsigned](samples []T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(samples))
	for i, v := range samples {
		if i == 0 {
			out = append(out, v)
			continue
		}
		out = append(out, v-samples[i-1])
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
