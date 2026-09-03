// Package tenantquota — Gopher Workplace challenge.
package tenantquota

import "sync"

// Quota tracks how many API units each tenant has consumed. Tenants appear
// and disappear at runtime, and requests for different tenants are counted
// concurrently.
type Quota struct {
	used sync.Map // tenant -> *counter
}

// Add charges n units to a tenant and returns the tenant's new total.
//
// Examples:
//
//	var q Quota; q.Add("acme", 3)             => 3
//	q.Add("acme", 2)                          => 5
//	q.Add("globex", 7)                        => 7
func (q *Quota) Add(tenant string, n int64) int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Used returns a tenant's total, or 0 if the tenant is unknown.
//
// Examples:
//
//	var q Quota; q.Used("nobody") => 0
func (q *Quota) Used(tenant string) int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Tenants returns every charged tenant, sorted.
//
// Examples:
//
//	q.Add("b", 1); q.Add("a", 1); q.Tenants() => ["a" "b"]
func (q *Quota) Tenants() []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
