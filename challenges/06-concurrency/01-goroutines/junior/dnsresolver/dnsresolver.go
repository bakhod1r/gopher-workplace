// Package dnsresolver — Gopher Workplace challenge.
package dnsresolver

// ResolveAll resolves every host, substituting a placeholder for failures.
//
// Examples:
//
//	ResolveAll([]string{"a.io", "b.io"}, lookup)  => [10.0.0.1 10.0.0.2]
//	ResolveAll([]string{"ghost.io"}, lookup)      => [0.0.0.0]
//	ResolveAll(nil, lookup)                       => []
func ResolveAll(hosts []string, resolve func(host string) string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
