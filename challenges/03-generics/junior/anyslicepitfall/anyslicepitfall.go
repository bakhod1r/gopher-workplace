// Package anyslicepitfall — Gopher Workplace challenge.
package anyslicepitfall

// ToAny converts a typed slice to []any, element by element.
// There is no way to do this without a copy.
func ToAny[T any](s []T) []any {
	// TODO(candidate): box each element into the result.
	panic("not implemented")
}

// SumInts totals a typed slice with no boxing at all.
// It is provided for comparison.
func SumInts(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
