// Package regionreplicator — Gopher Workplace challenge.
package regionreplicator

// Region is one geographic region and the availability zones inside it.
type Region struct {
	Name  string
	Zones []string
}

// ReplicateAll fans out one goroutine per region, and inside each region one
// goroutine per zone — a two-level tree. It returns the "region/zone" pairs that
// failed to replicate, sorted, so the runbook lists the same targets every time.
//
// Examples:
//
//	ReplicateAll([]Region{{"eu", []string{"a"}}}, replicate)          => []
//	ReplicateAll([]Region{{"eu", []string{"a", "bad"}}}, replicate)   => ["eu/bad"]
//	ReplicateAll(nil, replicate)                                      => []
func ReplicateAll(regions []Region, replicate func(region, zone string) error) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
