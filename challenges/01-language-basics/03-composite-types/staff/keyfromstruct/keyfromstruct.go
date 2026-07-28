// Package keyfromstruct counts coordinates using a struct map key. A planted bug
// swaps the fields when building the key.
package keyfromstruct

// Point is a comparable struct usable as a map key.
type Point struct {
	X, Y int
}

// Count tallies how many times each point appears.
func Count(pts []Point) map[Point]int {
	m := make(map[Point]int)
	for _, p := range pts {
		// CHANGE CODE BELOW THIS LINE
		m[Point{p.Y, p.X}]++
		// CHANGE CODE ABOVE THIS LINE
	}
	return m
}
