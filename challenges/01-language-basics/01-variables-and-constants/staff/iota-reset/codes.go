// Package codes numbers HTTP-like classes with iota. A planted explicit value
// breaks the running sequence.
package codes

// Class is a response class index.
type Class int

const (
	Info    Class = iota // 0
	Success              // 1
	// CHANGE CODE BELOW THIS LINE
	Redirect = 7
	// CHANGE CODE ABOVE THIS LINE
	ClientError // should be 3
	ServerError // should be 4
)
