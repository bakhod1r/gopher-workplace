// Package cleanupcount records how many resources are open at the moment a
// deferred cleanup runs. Because defers all fire at function return, every
// cleanup sees the FINAL open count — the test pins this behaviour, and a
// planted bug decrements at schedule time instead of via the defer.
package cleanupcount

// PeakThenDrain opens n resources (incrementing open), scheduling a deferred
// decrement for each, and returns the peak open count observed before draining.
func PeakThenDrain(n int) (peak int) {
	open := 0
	for i := 0; i < n; i++ {
		open++
		if open > peak {
			peak = open
		}
		// CHANGE CODE BELOW THIS LINE
		open--
		// CHANGE CODE ABOVE THIS LINE
	}
	return peak
}
