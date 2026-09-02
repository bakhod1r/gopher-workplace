// Package logshipper — Gopher Workplace challenge.
package logshipper

// PayloadSizes returns the wire size of every log line, including its newline.
//
// Examples:
//
//	PayloadSizes([]string{"ok", "boom"})  => [3 5]
//	PayloadSizes([]string{""})            => [1]
//	PayloadSizes(nil)                     => []
func PayloadSizes(lines []string) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
