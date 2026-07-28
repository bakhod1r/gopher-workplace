// Package mapdelnil removes a key from a map of pointers. A planted bug sets the
// value to nil but leaves the KEY present, so len and iteration still see it.
package mapdelnil

// Remove deletes id from the registry so it no longer appears.
func Remove(m map[int]*int, id int) {
	// CHANGE CODE BELOW THIS LINE
	m[id] = nil
	// CHANGE CODE ABOVE THIS LINE
}
