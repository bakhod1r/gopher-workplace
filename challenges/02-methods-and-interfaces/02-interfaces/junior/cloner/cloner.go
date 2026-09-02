// Package cloner — Gopher Workplace challenge.
package cloner

// Cloner is a thing that can duplicate itself.
type Cloner interface {
	Clone() Cloner
}

// Config is a named set of tags.
type Config struct {
	Name string
	Tags []string
}

// Clone returns an independent copy of the config.
//
// Examples:
//
//	(&Config{Name: "db"}).Clone().(*Config).Name => "db"
func (c *Config) Clone() Cloner {
	// TODO(candidate): copy the struct AND the Tags slice.
	panic("not implemented")
}

// CopyOf returns a duplicate of c.
func CopyOf(c Cloner) Cloner {
	// TODO(candidate): delegate to the value's own Clone.
	panic("not implemented")
}
