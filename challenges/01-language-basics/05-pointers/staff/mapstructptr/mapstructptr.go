// Package mapstructptr mutates an int referenced by a struct stored BY VALUE in
// a map. Because map struct values aren't addressable, a planted bug tries to
// mutate the map value's field directly is avoided; instead it fetches a copy
// and mutates the copy's non-pointer field. The pointer field still aliases, so
// mutate through the pointer.
package mapstructptr

type Ref struct{ P *int }

// BumpVia increments the int referenced by m[k].P. The Ref is stored by value,
// but its P pointer aliases a shared int.
func BumpVia(m map[int]Ref, k int) {
	r := m[k]
	// CHANGE CODE BELOW THIS LINE
	_ = r
	// CHANGE CODE ABOVE THIS LINE
}
