// Package deepclone — Gopher Workplace challenge.
package deepclone

// Config holds settings with a slice of tags.
type Config struct {
	Name string
	Tags []string
}

// Clone returns a deep copy of the Config. Modifying the clone's Tags must not
// affect the original.
//
// Examples:
//
//	c := Config{"app", []string{"v1", "prod"}}
//	c2 := c.Clone()
//	c2.Tags[0] = "v2" // c.Tags[0] is still "v1"
func (c Config) Clone() Config {
	// TODO(candidate): deep copy — don't share the Tags slice.
	panic("not implemented")
}
