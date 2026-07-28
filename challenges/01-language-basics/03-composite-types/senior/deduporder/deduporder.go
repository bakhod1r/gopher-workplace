// Package deduporder removes duplicates preserving first-seen order. A planted
// bug appends when the value was already seen.
package deduporder

// Unique returns xs with duplicates removed, keeping the first occurrence order.
func Unique(xs []int) []int {
	seen := make(map[int]struct{})
	out := []int{}
	for _, x := range xs {
		_, ok := seen[x]
		// CHANGE CODE BELOW THIS LINE
		if ok {
			// CHANGE CODE ABOVE THIS LINE
			out = append(out, x)
			seen[x] = struct{}{}
		}
	}
	return out
}
