// Package thresholdbug keeps elements strictly greater than a threshold. A
// planted bug uses >= instead of >, wrongly including the boundary value.
package thresholdbug

// AboveThreshold returns the elements of xs strictly greater than t.
func AboveThreshold(xs []int, t int) []int {
	var out []int
	for _, v := range xs {
		// CHANGE CODE BELOW THIS LINE
		if v >= t {
			// CHANGE CODE ABOVE THIS LINE
			out = append(out, v)
		}
	}
	return out
}
