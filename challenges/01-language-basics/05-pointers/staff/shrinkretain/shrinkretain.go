// Package shrinkretain shrinks a slice of pointers by one from the end. A planted
// bug only re-slices, leaving the dropped element's pointer live in the backing
// array (a retention leak). It must nil the dropped slot before shrinking.
package shrinkretain

// Pop removes the last element of a *[]*int, returning it, and must clear the
// vacated slot so the popped object can be garbage-collected.
func Pop(sp *[]*int) *int {
	s := *sp
	last := len(s) - 1
	v := s[last]
	// CHANGE CODE BELOW THIS LINE
	*sp = s[:last]
	// CHANGE CODE ABOVE THIS LINE
	return v
}
