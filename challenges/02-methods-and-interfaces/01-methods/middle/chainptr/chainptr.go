// Package chainptr — Gopher Workplace challenge.
package chainptr

// Config holds key-value settings.
type Config struct {
	Data map[string]string
}

// NewConfig creates an empty Config.
func NewConfig() *Config {
	return &Config{Data: make(map[string]string)}
}

// Set adds a key-value pair and returns the Config pointer for chaining.
//
// Examples:
//
//	c := NewConfig().Set("a", "1").Set("b", "2")
//	c.Data => {"a":"1", "b":"2"}
func (c *Config) Set(key, value string) *Config {
	// TODO(candidate): store the key-value pair and return c for chaining.
	panic("not implemented")
}
