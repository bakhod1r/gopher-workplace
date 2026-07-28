// Package deferargbug records the final value of a counter via defer. A planted
// bug passes the counter as a defer ARGUMENT, snapshotting it early, so the
// recorded value is the initial one, not the final.
package deferargbug

// FinalCount increments a counter n times and records the counter's value into
// a named result at function exit. It should record the FINAL value.
func FinalCount(n int) (recorded int) {
	c := 0
	// CHANGE CODE BELOW THIS LINE
	defer func(v int) { recorded = v }(c)
	// CHANGE CODE ABOVE THIS LINE
	for i := 0; i < n; i++ {
		c++
	}
	return
}
