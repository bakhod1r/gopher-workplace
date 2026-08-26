// Package mergeobj — Gopher Workplace challenge.
package mergeobj

// Config holds settings.
type Config struct {
	Host    string
	Port    int
	Debug   bool
}

// Merge applies non-zero fields from other into c.
// Zero-value fields in other are ignored.
//
// Examples:
//
//	c := Config{"localhost", 8080, false}
//	c.Merge(Config{Port: 9090, Debug: true})
//	// c == Config{"localhost", 9090, true}
func (c *Config) Merge(other Config) {
	// TODO(candidate): apply non-zero fields from other.
	panic("not implemented")
}
