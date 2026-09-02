// Package buildlog — Gopher Workplace challenge.
package buildlog

// TailReverse sends the stored build-log lines through a goroutine in
// newest-first order and returns what the renderer collected.
//
// Examples:
//
//	TailReverse([]string{"a","b","c"}) => ["c" "b" "a"]
//	TailReverse([]string{"x"}) => ["x"]
//	TailReverse(nil) => []
func TailReverse(lines []string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
