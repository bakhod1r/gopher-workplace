// Package pluginreg — Gopher Workplace challenge.
package pluginreg

import "errors"

// ErrDuplicate reports a repeated plugin name.
var ErrDuplicate = errors.New("duplicate plugin")

// Plugin is the required contract.
type Plugin interface {
	Name() string
	Run() string
}

// Closer is the optional cleanup contract.
type Closer interface {
	Close()
}

// Simple is a plugin with no cleanup.
type Simple struct {
	N string
}

// Name returns the plugin name.
func (s *Simple) Name() string { return s.N }

// Run returns the plugin's output.
func (s *Simple) Run() string { return "run:" + s.N }

// Closeable is a plugin that also cleans up.
type Closeable struct {
	N      string
	Closed bool
}

// Name returns the plugin name.
func (c *Closeable) Name() string { return c.N }

// Run returns the plugin's output.
func (c *Closeable) Run() string { return "run:" + c.N }

// Close marks the plugin closed.
func (c *Closeable) Close() { c.Closed = true }

// Registry holds plugins in registration order.
type Registry struct {
	plugins []Plugin
	names   map[string]bool
}

// NewRegistry returns an empty registry.
func NewRegistry() *Registry {
	return &Registry{names: make(map[string]bool)}
}

// Register adds a plugin, rejecting duplicate names.
func (r *Registry) Register(p Plugin) error {
	// TODO(candidate): reject duplicates, then store in order.
	panic("not implemented")
}

// RunAll runs every plugin in registration order.
func (r *Registry) RunAll() []string {
	// TODO(candidate): collect Run() output.
	panic("not implemented")
}

// CloseAll closes the plugins that implement Closer and returns the count.
//
// Examples:
//
//	one Closeable and one Simple => 1
func (r *Registry) CloseAll() int {
	// TODO(candidate): optional-interface assertion.
	panic("not implemented")
}
