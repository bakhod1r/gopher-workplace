// Package stackpush reverses via deferred pushes but a planted bug pushes with a
// body-captured slice index expression that reads the final loop state.
package stackpush

// ReverseInts returns xs reversed, built by deferring a push of each element.
func ReverseInts(xs []int) (out []int) {
	for i := 0; i < len(xs); i++ {
		// CHANGE CODE BELOW THIS LINE
		defer func() { out = append(out, xs[len(xs)-1]) }()
		// CHANGE CODE ABOVE THIS LINE
	}
	return
}
