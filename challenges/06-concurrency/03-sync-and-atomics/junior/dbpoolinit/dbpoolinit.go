// Package dbpoolinit - Gopher Workplace challenge.
package dbpoolinit

import "sync"

// Pool is a database connection pool handle.
type Pool struct {
	DSN string
}

// Provider hands out a lazily opened connection pool.
type Provider struct {
	once  sync.Once
	open  func() *Pool
	pool  *Pool
	opens int
}

// NewProvider returns a Provider that calls open on first use.
func NewProvider(open func() *Pool) *Provider {
	return &Provider{open: open}
}

// Pool returns the connection pool, opening it on the first call only.
//
// Examples:
//
//	p := NewProvider(func() *Pool { return &Pool{DSN: "db"} }); p.Pool().DSN => "db"
//	p.Pool(); p.Pool()                                                       => same pool, opened once
func (p *Provider) Pool() *Pool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Opens reports how many times the open function actually ran.
//
// Examples:
//
//	p.Pool(); p.Pool(); p.Opens() => 1
func (p *Provider) Opens() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
