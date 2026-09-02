// Package healthcheck — Gopher Workplace challenge.
package healthcheck

// CountHealthy probes every host concurrently and returns how many of them
// reported healthy, using a mutex to protect the shared counter.
//
// Examples:
//
//	CountHealthy([]string{"ok-a", "bad"}, isOK)  => 1
//	CountHealthy([]string{"ok-a", "ok-b"}, isOK) => 2
//	CountHealthy(nil, isOK)                      => 0
func CountHealthy(hosts []string, check func(string) bool) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
