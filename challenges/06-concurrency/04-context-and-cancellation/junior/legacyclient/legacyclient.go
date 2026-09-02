// Package legacyclient — Gopher Workplace challenge.
package legacyclient

import "context"

// LegacyClientContext returns the context the payment-gateway SDK wrapper uses
// until the surrounding call chain is migrated to accept a real request
// context. It marks the call site as unfinished plumbing rather than claiming
// to be a legitimate root.
//
// Examples:
//
//	LegacyClientContext().Err()               => nil
//	LegacyClientContext().Done()              => nil
//	LegacyClientContext() == context.TODO()   => true
func LegacyClientContext() context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}
