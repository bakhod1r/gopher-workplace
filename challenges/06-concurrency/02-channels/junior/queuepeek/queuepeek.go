// Package queuepeek — Gopher Workplace challenge.
package queuepeek

// PeekJob attempts one non-blocking receive on the job queue using select
// with default. It returns the job id and true if a receive could proceed
// immediately, otherwise 0 and false.
//
// A closed queue is always ready to receive.
//
// Examples:
//
//	PeekJob(chan with 5)     => 5, true
//	PeekJob(empty open queue) => 0, false
//	PeekJob(closed queue)     => 0, true
func PeekJob(jobs <-chan int) (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
