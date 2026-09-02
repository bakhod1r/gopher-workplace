// Package healthpings — Gopher Workplace challenge.
package healthpings

// HealthPings emits count probe URLs for the given endpoint on a fresh
// channel and closes it. A count of zero or less emits nothing.
//
// Examples:
//
//	HealthPings("api", 2)   => yields "api/health" twice then closes
//	HealthPings("db", 1)    => yields "db/health" then closes
//	HealthPings("api", 0)   => closes immediately
func HealthPings(endpoint string, count int) <-chan string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
