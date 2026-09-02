// Package logger — Gopher Workplace challenge.
package logger

// Logger records one line of output.
type Logger interface {
	Log(line string)
}

// MemLogger keeps every line in memory.
type MemLogger struct {
	Lines []string
}

// Log appends the line.
//
// Examples:
//
//	m := &MemLogger{}; m.Log("a") => m.Lines == ["a"]
func (m *MemLogger) Log(line string) {
	// TODO(candidate): append to m.Lines.
	panic("not implemented")
}

// Discard throws every line away.
type Discard struct{}

// Log ignores the line.
func (d Discard) Log(line string) {
	// TODO(candidate): do nothing.
	panic("not implemented")
}

// LogAll writes every line, in order.
func LogAll(l Logger, lines []string) {
	// TODO(candidate): log each line.
	panic("not implemented")
}
