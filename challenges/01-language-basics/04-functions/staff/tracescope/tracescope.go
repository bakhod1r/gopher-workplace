// Package tracescope logs a start and end marker per item; each item's "end"
// must be logged before the next item's "start". A planted bug uses a
// FUNCTION-level defer inside the loop, so all "end" markers run at function
// return (in LIFO order) instead of at the end of each iteration.
package tracescope

import "fmt"

// Trace returns the interleaved start/end log. For n=2 it must be
// [start0 end0 start1 end1].
func Trace(n int) (log []string) {
	for i := 0; i < n; i++ {
		// CHANGE CODE BELOW THIS LINE
		log = append(log, fmt.Sprintf("start%d", i))
		defer func(k int) { log = append(log, fmt.Sprintf("end%d", k)) }(i)
		// CHANGE CODE ABOVE THIS LINE
	}
	return
}
