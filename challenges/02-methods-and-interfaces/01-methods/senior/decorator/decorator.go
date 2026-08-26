// Package decorator — Gopher Workplace challenge.
package decorator

// Component does work.
type Component struct{}

// DoWork returns a string.
func (c *Component) DoWork() string {
	return "work"
}

// Decorator wraps a component.
type Decorator struct {
	Comp *Component
}

// DoWork calls the wrapped component's DoWork and surrounds it with brackets: "[work]".
func (d *Decorator) DoWork() string {
	// TODO(candidate): call d.Comp.DoWork() and wrap with brackets.
	panic("not implemented")
}
