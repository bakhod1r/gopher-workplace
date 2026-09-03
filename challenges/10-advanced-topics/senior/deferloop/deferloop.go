// Package deferloop — Gopher Workplace challenge.
package deferloop

// Process doubles each item and calls release with the item as soon as
// that item is finished.
//
// release returns the item's resources. Holding every item until the
// function returns is what makes a batch job run out of them.
//
// Examples:
//
//	Process([]int{1, 2}, rel) => []int{2, 4}, rel called after each item
func Process(items []int, release func(int)) []int {
	// CHANGE CODE BELOW THIS LINE
	out := make([]int, 0, len(items))
	for _, v := range items {
		defer release(v)
		out = append(out, v*2)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
