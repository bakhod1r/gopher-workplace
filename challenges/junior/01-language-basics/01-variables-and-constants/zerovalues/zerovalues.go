// Package zerovalues — Gopher Workplace challenge.
package zerovalues

// Config holds server settings. Any field the caller leaves unset must keep
// Go's zero value for its type.
type Config struct {
	Host  string
	Port  int
	Debug bool
	Tags  []string
}

// DefaultConfig returns the baseline Config: only Port is set, to 8080.
// Every other field is left at its type's zero value — Host "", Debug false,
// Tags nil (not an empty non-nil slice).
//
// Tip: naming the port as a constant (`const DefaultPort = 8080`) instead of a
// bare literal is good style — but only the returned values are graded.
//
// Examples:
//
//	DefaultConfig().Port  => 8080
//	DefaultConfig().Host  => ""
//	DefaultConfig().Tags  => nil
func DefaultConfig() Config {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
