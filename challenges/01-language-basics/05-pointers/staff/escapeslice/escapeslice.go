// Package escapeslice builds a slice of pointers to per-element structs. A
// planted bug appends the address of a single hoisted struct reused across
// iterations, so all pointers alias the final value.
package escapeslice

type Item struct{ V int }

// Items returns one *Item per input value; each pointer must hold a DISTINCT V.
func Items(vs []int) []*Item {
	var out []*Item
	var it Item
	for _, v := range vs {
		it.V = v
		// CHANGE CODE BELOW THIS LINE
		out = append(out, &it)
		// CHANGE CODE ABOVE THIS LINE
	}
	return out
}
