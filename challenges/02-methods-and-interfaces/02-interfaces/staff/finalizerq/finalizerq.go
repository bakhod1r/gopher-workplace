// Package finalizerq — Gopher Workplace challenge.
package finalizerq

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// Resource is the underlying releasable thing.
type Resource interface {
	Release()
}

// Pool counts releases; it stands in for an OS resource pool.
type Pool struct {
	released atomic.Int64
}

// Release records one release.
func (p *Pool) Release() { p.released.Add(1) }

// Released returns how many releases happened.
func (p *Pool) Released() int64 { return p.released.Load() }

// Handle owns a resource and releases it at most once.
type Handle struct {
	pool   Resource
	once   sync.Once
	closed atomic.Bool
}

// NewHandle returns a handle with a finalizer registered as a safety net.
//
// Examples:
//
//	a dropped handle => released by the finalizer after a GC
func NewHandle(p Resource) *Handle {
	// TODO(candidate): build the handle, register the finalizer.
	panic("not implemented")
}

// Close releases the resource and clears the finalizer. It is idempotent.
func (h *Handle) Close() {
	// TODO(candidate): release once, then clear the finalizer.
	panic("not implemented")
}

// Closed reports whether the resource has been released.
func (h *Handle) Closed() bool {
	// TODO(candidate): report the state.
	panic("not implemented")
}

var _ = runtime.SetFinalizer
