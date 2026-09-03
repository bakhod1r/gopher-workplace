// Package boxvalues — Gopher Workplace challenge.
package boxvalues

// Total sums vals.
//
// Passing the values through []any boxes every element: an interface value
// needs a word to point at, so each int gets a heap home it never needed.
//
// Examples:
//
//	Total([]int{1, 2, 3}) => 6
func Total(vals []int) int64 {
	// CHANGE CODE BELOW THIS LINE
	boxed := make([]any, 0, len(vals))
	for _, v := range vals {
		boxed = append(boxed, v)
	}
	var total int64
	for _, b := range boxed {
		total += int64(b.(int))
	}
	return total
	// CHANGE CODE ABOVE THIS LINE
}
