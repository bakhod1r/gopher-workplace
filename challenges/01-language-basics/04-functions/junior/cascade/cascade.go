// Package cascade — Gopher Workplace challenge.
package cascade

// Access returns the comma-separated permissions granted at an access level,
// highest first, using switch fallthrough so each higher level accumulates the
// lower ones:
//
//	3 => "admin,write,read"
//	2 => "write,read"
//	1 => "read"
//	anything else => ""
//
// Examples:
//
//	Access(3) => "admin,write,read"
//	Access(2) => "write,read"
//	Access(1) => "read"
//	Access(0) => ""
func Access(level int) string {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
