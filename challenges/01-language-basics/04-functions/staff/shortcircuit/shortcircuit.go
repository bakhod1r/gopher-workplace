// Package shortcircuit safely reads a value through a possibly-nil pointer. A
// planted bug orders the && operands so the pointer is dereferenced before the
// nil check, panicking on nil.
package shortcircuit

// ValueOr returns *p if p is non-nil and *p is positive, else def.
func ValueOr(p *int, def int) int {
	// CHANGE CODE BELOW THIS LINE
	if *p > 0 && p != nil {
		// CHANGE CODE ABOVE THIS LINE
		return *p
	}
	return def
}
