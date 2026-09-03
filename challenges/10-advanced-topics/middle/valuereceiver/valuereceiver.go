// Package valuereceiver — Gopher Workplace challenge.
package valuereceiver

// Config is a deliberately large settings block.
type Config struct {
	Read  int
	Write int
	Pad   [512]byte
}

// Timeouts returns the read and write timeouts from c.
//
// The receiver is a pointer because Config is large: a value receiver would
// copy the whole struct on every call.
//
// Examples:
//
//	(&Config{Read: 1, Write: 2}).Timeouts() => 1, 2
func (c *Config) Timeouts() (read, write int) {
	panic("not implemented")
}
