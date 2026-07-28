// Package appendtwobranch appends two different tails to a base slice. A planted
// bug lets both appends share the base's spare capacity, so the second clobbers
// the first.
package appendtwobranch

// Branch returns two independent slices: base+x and base+y.
func Branch(a []int, x, y int) ([]int, []int) {
	// CHANGE CODE BELOW THIS LINE
	base := a
	// CHANGE CODE ABOVE THIS LINE
	b := append(base, x)
	c := append(base, y)
	return b, c
}
