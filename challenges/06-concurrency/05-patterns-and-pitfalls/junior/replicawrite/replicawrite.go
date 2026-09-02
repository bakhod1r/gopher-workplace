// Package replicawrite — Gopher Workplace challenge.
package replicawrite

// FirstReplicaAck writes to every replica at once and reports whether any of
// them acknowledged. It returns as soon as an ack arrives, without leaking
// the goroutines that are still writing.
//
// Examples:
//
//	FirstReplicaAck([]string{"ok-1", "bad-2"}, write)  => true
//	FirstReplicaAck([]string{"bad-1"}, write)          => false
//	FirstReplicaAck(nil, write)                        => false
func FirstReplicaAck(replicas []string, write func(string) bool) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
