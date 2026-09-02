// Package healthprobe — Gopher Workplace challenge.
package healthprobe

// Probe runs rounds probe/ack round trips between two goroutines over
// unbuffered channels and returns the trace: "probe", "ack", ...
//
// Examples:
//
//	Probe(1)  => ["probe" "ack"]
//	Probe(2)  => ["probe" "ack" "probe" "ack"]
//	Probe(0)  => []
func Probe(rounds int) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
