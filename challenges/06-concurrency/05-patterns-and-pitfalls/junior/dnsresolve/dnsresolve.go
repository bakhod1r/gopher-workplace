// Package dnsresolve — Gopher Workplace challenge.
package dnsresolve

// ResolveAll resolves every host concurrently and returns a host to address
// map, with the shared map guarded by a mutex.
//
// Examples:
//
//	ResolveAll([]string{"a"}, lookup)       => map[a:10.0.0.1]
//	ResolveAll([]string{"a", "b"}, lookup)  => map with both hosts
//	ResolveAll(nil, lookup)                 => empty map
func ResolveAll(hosts []string, lookup func(string) string) map[string]string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
