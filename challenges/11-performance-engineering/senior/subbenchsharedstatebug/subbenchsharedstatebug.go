// Package subbenchsharedstatebug — Gopher Workplace challenge.
package subbenchsharedstatebug

// Runner drives one sub-benchmark per input size, accumulating each size's
// result independently — the b.Run("size=N", ...) pattern.
type Runner struct {
	buf []int
}

// RunSize processes size items and returns their sum, using the runner's
// scratch buffer. Each sub-benchmark must start from a clean buffer:
// carrying the previous size's data over silently inflates every result
// after the first.
//
// Examples:
//
//	r.RunSize(3) => 3 (1+1+1)
func (r *Runner) RunSize(size int) int {
	// CHANGE CODE BELOW THIS LINE
	// (the buffer is reused across sizes)
	// CHANGE CODE ABOVE THIS LINE
	for i := 0; i < size; i++ {
		r.buf = append(r.buf, 1)
	}
	sum := 0
	for _, v := range r.buf {
		sum += v
	}
	return sum
}

// RunAll runs one sub-benchmark per size and returns the results in order.
//
// Examples:
//
//	RunAll([]int{1, 2, 3}) => []int{1, 2, 3}
func RunAll(sizes []int) []int {
	var r Runner
	out := make([]int, 0, len(sizes))
	for _, s := range sizes {
		out = append(out, r.RunSize(s))
	}
	return out
}
