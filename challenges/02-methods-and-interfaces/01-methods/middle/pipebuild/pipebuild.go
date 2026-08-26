// Package pipebuild — Gopher Workplace challenge.
package pipebuild

import "strings"

// Pipe is a string processor.
type Pipe struct {
	text string
}

// NewPipe starts a new pipeline.
func NewPipe(t string) *Pipe {
	return &Pipe{text: t}
}

// Upper converts text to uppercase and returns the Pipe for chaining.
func (p *Pipe) Upper() *Pipe {
	// TODO(candidate): use strings.ToUpper
	_ = strings.ToUpper
	panic("not implemented")
}

// Replace replaces old with new in the text and returns the Pipe.
func (p *Pipe) Replace(old, new string) *Pipe {
	// TODO(candidate): use strings.ReplaceAll
	panic("not implemented")
}

// Result returns the final string.
func (p *Pipe) Result() string {
	return p.text
}
