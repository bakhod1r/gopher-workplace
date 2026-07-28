// Package handlerbug builds label-printer closures from a slice. A planted bug
// captures the shared index variable `i` (declared once, outside the range) so
// every closure reads the same final index.
package handlerbug

// Labelers returns one closure per element; the k-th returns names[k].
func Labelers(names []string) []func() string {
	var out []func() string
	// CHANGE CODE BELOW THIS LINE
	var i int
	for i = 0; i < len(names); i++ {
		out = append(out, func() string { return names[i] })
	}
	// CHANGE CODE ABOVE THIS LINE
	return out
}
