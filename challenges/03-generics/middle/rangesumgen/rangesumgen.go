// Package rangesumgen — Gopher Workplace challenge.
package rangesumgen

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// SumRange returns lo + (lo+1) + ... + hi.
// It returns 0 when hi < lo.
func SumRange[T Integer](lo, hi T) T {
	// TODO(candidate): use the closed-form sum, not a loop.
	panic("not implemented")
}
