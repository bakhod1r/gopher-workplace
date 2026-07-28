// Package iifeinit initialises a lookup table via an immediately-invoked
// function. A planted bug assigns the function VALUE instead of invoking it, so
// the table field holds a func, and the builder returns an empty table.
package iifeinit

// BuildTable returns a map from i to i*i for i in [0,n), built via an
// immediately-invoked function literal.
func BuildTable(n int) map[int]int {
	// CHANGE CODE BELOW THIS LINE
	table := func() map[int]int {
		m := map[int]int{}
		for i := 0; i < n; i++ {
			m[i] = i * i
		}
		return m
	}
	_ = table
	return nil
	// CHANGE CODE ABOVE THIS LINE
}
