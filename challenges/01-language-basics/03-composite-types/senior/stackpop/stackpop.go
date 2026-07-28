// Package stackpop pops the top of a stack. A planted bug returns the element
// but forgets to shrink the stack.
package stackpop

// Pop removes and returns the last element of s, returning the shrunken stack
// and the popped value. Popping an empty stack returns (s, 0, false).
func Pop(s []int) ([]int, int, bool) {
	if len(s) == 0 {
		return s, 0, false
	}
	top := s[len(s)-1]
	// CHANGE CODE BELOW THIS LINE
	return s, top, true
	// CHANGE CODE ABOVE THIS LINE
}
