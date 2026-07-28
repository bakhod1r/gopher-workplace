// Package boundmethod — Gopher Workplace challenge.
package boundmethod

type Counter struct{ N int }

func (c *Counter) Inc() { c.N++ }

// Bind returns a function that, when called, increments c. It captures the
// pointer so all invocations affect the same counter.
func Bind(c *Counter) func() {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
