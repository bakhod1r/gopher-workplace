// Package typesetmethodgen — Gopher Workplace challenge.
package typesetmethodgen

// Code is a named int type that can label itself.
type Code interface {
	~int
	String() string
}

// Status is an HTTP-like status code.
type Status int

// String labels the status.
func (s Status) String() string {
	switch s {
	case 200:
		return "ok"
	case 404:
		return "missing"
	default:
		return "other"
	}
}

// Labels returns the String() of every code, and their numeric sum.
func Labels[T Code](s []T) ([]string, int) {
	// TODO(candidate): collect the labels and add up the numeric values.
	panic("not implemented")
}
