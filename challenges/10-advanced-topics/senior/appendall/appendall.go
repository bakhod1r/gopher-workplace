// Package appendall — Gopher Workplace challenge.
package appendall

// AppendAll returns every part concatenated in order.
//
// The final length is known before the first append, so the result should
// be allocated once instead of growing through every doubling.
//
// Examples:
//
//	AppendAll([][]int{{1}, {2, 3}}) => []int{1, 2, 3}
func AppendAll(parts [][]int) []int {
	// CHANGE CODE BELOW THIS LINE
	var out []int
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
