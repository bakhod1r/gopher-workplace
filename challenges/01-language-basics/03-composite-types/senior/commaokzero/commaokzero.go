// Package commaokzero looks up config with a default.
// A planted bug treats a stored zero as "missing".
package commaokzero

// GetOr returns m[key] if the key is present, else def.
func GetOr(m map[string]int, key string, def int) int {
	// CHANGE CODE BELOW THIS LINE
	if v := m[key]; v != 0 {
		return v
	}
	// CHANGE CODE ABOVE THIS LINE
	return def
}
