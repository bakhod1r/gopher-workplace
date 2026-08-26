// Package applyfn — Gopher Workplace challenge.
package applyfn

// Transformer holds a multiplier.
type Transformer struct {
	Factor int
}

// Transform returns n * Factor.
func (t Transformer) Transform(n int) int {
	return n * t.Factor
}

// ApplyAll applies fn to each element of nums and returns the results.
//
// Examples:
//
//	t := Transformer{Factor: 2}
//	ApplyAll(t.Transform, []int{1, 2, 3}) => [2, 4, 6]
func ApplyAll(fn func(int) int, nums []int) []int {
	// TODO(candidate): apply fn to each element, collect results.
	panic("not implemented")
}
