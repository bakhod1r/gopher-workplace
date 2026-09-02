// Package ifacemap — Gopher Workplace challenge.
package ifacemap

import "errors"

// ErrUnknown is returned for an unregistered command name.
var ErrUnknown = errors.New("unknown command")

// Command transforms an argument.
type Command interface {
	Exec(arg string) string
}

// CommandFunc adapts a function to Command.
type CommandFunc func(string) string

// Exec calls the underlying function.
func (f CommandFunc) Exec(arg string) string { return f(arg) }

// Registry dispatches by name.
type Registry struct {
	cmds map[string]Command
}

// NewRegistry returns an empty registry.
func NewRegistry() *Registry {
	return &Registry{cmds: make(map[string]Command)}
}

// Register adds or replaces a command.
func (r *Registry) Register(name string, c Command) {
	// TODO(candidate): store it under name.
	panic("not implemented")
}

// Run executes the named command.
//
// Examples:
//
//	r.Run("up", "hi")   => "HI", nil
//	r.Run("nope", "hi") => "", ErrUnknown
func (r *Registry) Run(name, arg string) (string, error) {
	// TODO(candidate): look up, then execute.
	panic("not implemented")
}

// Names returns the registered names in sorted order.
func (r *Registry) Names() []string {
	// TODO(candidate): collect the keys and sort them.
	panic("not implemented")
}
