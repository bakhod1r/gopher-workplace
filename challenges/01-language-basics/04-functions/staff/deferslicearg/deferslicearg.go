// Package deferslicearg logs the FINAL length of a slice via defer. A planted
// bug passes the slice as a deferred argument, snapshotting its header (length 0)
// before elements are appended, so the logged length is wrong.
package deferslicearg

// BuildAndReport appends n items to a slice and returns the length recorded by a
// deferred reporter. The reporter must see the FINAL length.
func BuildAndReport(n int) (reported int) {
	var xs []int
	// CHANGE CODE BELOW THIS LINE
	defer func(s []int) { reported = len(s) }(xs)
	// CHANGE CODE ABOVE THIS LINE
	for i := 0; i < n; i++ {
		xs = append(xs, i)
	}
	return
}
