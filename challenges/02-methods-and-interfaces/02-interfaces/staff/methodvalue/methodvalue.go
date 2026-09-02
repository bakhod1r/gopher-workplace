// Package methodvalue — Gopher Workplace challenge.
package methodvalue

// Getter reads a value.
type Getter interface {
	Get() int
}

// Counter is mutated through a pointer.
type Counter struct {
	N int
}

// Get returns the current value.
func (c *Counter) Get() int {
	// TODO(candidate): return N.
	panic("not implemented")
}

// Set replaces the value.
func (c *Counter) Set(n int) {
	// TODO(candidate): store N.
	panic("not implemented")
}

// ValCounter is read through a value receiver.
type ValCounter struct {
	N int
}

// Get returns the value of this copy.
func (v ValCounter) Get() int {
	// TODO(candidate): return N.
	panic("not implemented")
}

// BindValue returns the method value v.Get, which captures a copy of v.
//
// Examples:
//
//	bind, then mutate v => the closure still returns the old value
func BindValue(v ValCounter) func() int {
	// TODO(candidate): the method value bound to the value receiver.
	panic("not implemented")
}

// BindPointer returns the method value c.Get, which captures the pointer.
func BindPointer(c *Counter) func() int {
	// TODO(candidate): the method value bound to the pointer receiver.
	panic("not implemented")
}

// GetExpr is the method expression form: the receiver is an argument.
func GetExpr() func(ValCounter) int {
	// TODO(candidate): the method expression ValCounter.Get.
	panic("not implemented")
}
