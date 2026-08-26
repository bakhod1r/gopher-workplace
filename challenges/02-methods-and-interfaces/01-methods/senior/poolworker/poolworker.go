// Package poolworker — Gopher Workplace challenge.
package poolworker

import "sync"

// Pool manages workers.
type Pool struct {
	Count int
	Tasks chan func()
	wg    sync.WaitGroup
}

// Start launches Count workers. Each worker reads from Tasks and executes them
// until Tasks is closed.
func (p *Pool) Start() {
	// TODO(candidate): start p.Count goroutines. Each loops over p.Tasks.
	// Use p.wg.Add(1) before starting, and p.wg.Done() inside the worker.
	panic("not implemented")
}

// Wait blocks until all workers finish (Tasks must be closed first).
func (p *Pool) Wait() {
	p.wg.Wait()
}
