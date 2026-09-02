# Lazy Database Pool

## Intuition

`sync.Once` is a one-shot gate. The first goroutine through runs the function; everyone else blocks until it returns, then proceeds. That gives both mutual exclusion and a happens-before edge, so the cached pointer is safely visible to every later caller.

## Approach

1. Store `once sync.Once`, `pool *Pool`, `opens int`.
2. In `Pool`, call `p.once.Do(...)` with a closure that opens the pool and records it.
3. Return the cached field afterwards.

## Solution

```go
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
	p.once.Do(func() {
		p.pool = p.open()
		p.opens++
	})
	return p.pool
}

// Opens reports how many times the open function actually ran.
//
// Examples:
//
//	p.Pool(); p.Pool(); p.Opens() => 1
func (p *Provider) Opens() int {
	return p.opens
}
```

## Walkthrough

Ten request goroutines call `Pool` at start-up. One wins the `Do`, opens the pool and sets `opens = 1`. The other nine block until it returns, then all read the same pointer.

## Pitfalls

- Calling the open function outside `once.Do` - it then runs on every request.
- Reusing a `sync.Once` expecting it to fire again after a reconnect; it never does.
- Copying the provider after first use, which copies the `Once`.
