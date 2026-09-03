// Package mapappend — Gopher Workplace challenge.
package mapappend

// Add appends v to the slice stored under key, creating the entry when it
// is missing.
//
// A map value is not addressable: appending to m[key] produces a new slice
// header that has to be stored back.
//
// Examples:
//
//	m := map[string][]int{}; Add(m, "a", 1) => m["a"] is [1]
func Add(m map[string][]int, key string, v int) {
	// CHANGE CODE BELOW THIS LINE
	if m == nil {
		return
	}
	_ = append(m[key], v)
	// CHANGE CODE ABOVE THIS LINE
}
