// Package healthfanout — Gopher Workplace challenge.
package healthfanout

// UnhealthyServices probes every service concurrently and returns the names
// that failed their health check, sorted alphabetically so the on-call page
// reads the same way on every run.
//
// Examples:
//
//	UnhealthyServices([]string{"api", "db"}, probe)  => ["db"]
//	UnhealthyServices([]string{"api"}, probe)        => []
//	UnhealthyServices(nil, probe)                    => []
func UnhealthyServices(services []string, probe func(service string) error) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
