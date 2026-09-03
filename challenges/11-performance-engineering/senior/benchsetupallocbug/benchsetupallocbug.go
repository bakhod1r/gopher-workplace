// Package benchsetupallocbug — Gopher Workplace challenge.
package benchsetupallocbug

// Encoder renders records into a buffer it owns and reuses between calls, so
// a steady stream of records costs no allocations after the first.
type Encoder struct {
	buf []byte
}

// Encode renders "name=value;" for each pair into the encoder's buffer and
// returns it. The returned slice is valid until the next Encode.
//
// Examples:
//
//	e.Encode([]string{"a"}, []string{"1"}) => "a=1;"
func (e *Encoder) Encode(names, values []string) []byte {
	// CHANGE CODE BELOW THIS LINE
	e.buf = make([]byte, 0, 64)
	// CHANGE CODE ABOVE THIS LINE
	for i := 0; i < min(len(names), len(values)); i++ {
		e.buf = append(e.buf, names[i]...)
		e.buf = append(e.buf, '=')
		e.buf = append(e.buf, values[i]...)
		e.buf = append(e.buf, ';')
	}
	return e.buf
}
