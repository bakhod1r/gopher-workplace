// Package semverparse parses a "major.minor.patch" version string.
package semverparse

// Parse splits a version like "1.4.10" into its three integer components.
// Returns (0,0,0,false) if the format is not exactly three dot-separated
// non-negative integers.
//
// TODO(candidate): implement without strings.Split or strconv — scan bytes.
func Parse(s string) (major, minor, patch int, ok bool) {
	panic("not implemented")
}
