// Package hoistedaddr collects pointers to each value. A planted bug takes the
// address of a single hoisted variable reused across iterations, so every
// pointer ends up aliasing the same (final) value.
package hoistedaddr

// Pointers returns a pointer to a copy of each element. Each pointer must hold a
// DISTINCT value.
func Pointers(xs []int) []*int {
	var out []*int
	var v int
	for i := 0; i < len(xs); i++ {
		v = xs[i]
		// CHANGE CODE BELOW THIS LINE
		out = append(out, &v)
		// CHANGE CODE ABOVE THIS LINE
	}
	return out
}
