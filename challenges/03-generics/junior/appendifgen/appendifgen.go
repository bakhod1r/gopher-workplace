// Package appendifgen — Gopher Workplace challenge.
package appendifgen

// AppendIf appends v to s when cond is true, and returns s unchanged
// otherwise.
//
// Examples:
//
//	AppendIf([]int{1}, 2, true)  => []int{1, 2}
//	AppendIf([]int{1}, 2, false) => []int{1}
func AppendIf[T any](s []T, v T, cond bool) []T {
	// TODO(candidate): append only when cond is true.
	panic("not implemented")
}
