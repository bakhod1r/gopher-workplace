// Package strategypatt — Gopher Workplace challenge.
package strategypatt

// Context holds data to process.
type Context struct {
	Data []int
}

// Process applies the given strategy function to all elements in place.
func (c *Context) Process(strategy func(int) int) {
	// TODO(candidate): apply strategy to each element.
	panic("not implemented")
}
