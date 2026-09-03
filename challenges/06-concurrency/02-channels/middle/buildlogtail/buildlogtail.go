// Package buildlogtail — Gopher Workplace challenge.
package buildlogtail

// TailBuildLog follows a build's log stream and returns only the last keep
// lines, in stream order. The log is unbounded and the CI page only renders a
// tail, so memory is capped by keep: a buffered channel of capacity keep acts
// as the ring, and the oldest line is discarded whenever a new one arrives at
// a full ring.
//
// Examples:
//
//	TailBuildLog(chan a,b,c, 2) => [b c]
//	TailBuildLog(chan a,b, 5)   => [a b]
//	TailBuildLog(chan a,b, 0)   => []
func TailBuildLog(lines <-chan string, keep int) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
