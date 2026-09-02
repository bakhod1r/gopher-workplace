// Package statuscount — Gopher Workplace challenge.
package statuscount

// CountStatus drains the log line stream and returns how many lines carry
// the wanted status code.
//
// Examples:
//
//	CountStatus(chan "200","500","200", "200") => 2
//	CountStatus(chan "200", "404") => 0
//	CountStatus(closed empty, "200") => 0
func CountStatus(lines <-chan string, want string) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
