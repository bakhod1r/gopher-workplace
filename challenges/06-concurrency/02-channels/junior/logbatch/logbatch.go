// Package logbatch — Gopher Workplace challenge.
package logbatch

// CollectLines drains lines until the channel is closed and returns the
// batch in arrival order. It always returns a non-nil slice.
//
// Examples:
//
//	CollectLines(chan "a","b") => ["a" "b"]
//	CollectLines(closed empty) => []
//	CollectLines(chan "x")     => ["x"]
func CollectLines(lines <-chan string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
