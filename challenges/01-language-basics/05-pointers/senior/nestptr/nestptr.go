// Package nestptr records counts per (group, key). A planted bug writes to an
// inner map that was never created, panicking on the first insert.
package nestptr

// Add increments m[group][key], creating the inner map when needed.
func Add(m map[string]map[string]int, group, key string) {
	// CHANGE CODE BELOW THIS LINE
	m[group][key]++
	// CHANGE CODE ABOVE THIS LINE
}
