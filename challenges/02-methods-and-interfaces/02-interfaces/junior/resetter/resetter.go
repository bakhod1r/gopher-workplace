// Package resetter — Gopher Workplace challenge.
package resetter

// Resetter returns itself to its zero state.
type Resetter interface {
	Reset()
}

// Buffer holds strings.
type Buffer struct {
	Data []string
}

// Reset empties the buffer.
func (b *Buffer) Reset() {
	// TODO(candidate): drop the data.
	panic("not implemented")
}

// Gauge holds a number.
type Gauge struct {
	Value int
}

// Reset zeroes the gauge.
func (g *Gauge) Reset() {
	// TODO(candidate): back to zero.
	panic("not implemented")
}

// ResetAll resets every element.
func ResetAll(rs []Resetter) {
	// TODO(candidate): reset each one.
	panic("not implemented")
}
