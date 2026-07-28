// Package arraysharing returns two slice views over the same array and expects
// them to alias. A planted bug copies the array into a fresh slice for the second
// view, so writes through one are not seen by the other.
package arraysharing

// Views returns two slices that both view the array p points to (they must
// alias: a write through one is visible in the other).
func Views(p *[3]int) ([]int, []int) {
	a := p[:]
	// CHANGE CODE BELOW THIS LINE
	b := append([]int(nil), p[:]...)
	// CHANGE CODE ABOVE THIS LINE
	return a, b
}
