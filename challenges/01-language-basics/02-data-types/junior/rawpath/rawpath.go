// Package rawpath returns a Windows-style path using a raw string literal.
package rawpath

// TempPath returns the literal string C:\Users\temp\log.txt using a raw string
// literal (backticks), so no backslash needs escaping.
//
// TODO(candidate): return the path as a raw string literal.
func TempPath() string {
	panic("not implemented")
}
