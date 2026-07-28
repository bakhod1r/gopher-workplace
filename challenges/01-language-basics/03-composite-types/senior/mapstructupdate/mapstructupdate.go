// Package mapstructupdate updates a struct stored in a map. A planted bug mutates
// a copy and never writes it back.
package mapstructupdate

// Stat tracks hits.
type Stat struct {
	Hits int
}

// Record increments the Hits of key in m (creating it if absent).
func Record(m map[string]Stat, key string) {
	s := m[key]
	s.Hits++
	// CHANGE CODE BELOW THIS LINE
	_ = s
	// CHANGE CODE ABOVE THIS LINE
}
