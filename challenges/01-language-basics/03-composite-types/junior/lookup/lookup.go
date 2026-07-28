// Package lookup — Gopher Workplace challenge.
package lookup

// Lookup reports the value stored for key in m and whether the key was present.
// It must use the comma-ok form of a map read, so a key that is present but
// maps to 0 is still reported as found — a plain `m[key]` cannot tell "absent"
// from "present with the zero value".
//
// Examples:
//
//	Lookup(map[string]int{"a": 5}, "a") => 5, true
//	Lookup(map[string]int{"a": 5}, "z") => 0, false
//	Lookup(map[string]int{"z": 0}, "z") => 0, true   // present, value is 0
//	Lookup(nil, "a")                    => 0, false
func Lookup(m map[string]int, key string) (int, bool) {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
