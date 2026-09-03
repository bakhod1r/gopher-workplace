// Package deepequalcost — Gopher Workplace challenge.
package deepequalcost

import "reflect"

// Config is a comparable settings block.
type Config struct {
	Retries int
	Timeout int
	Name    string
	Debug   bool
}

// Changed reports whether the two configs differ.
//
// Config is a comparable struct, so == does the whole job. Reflecting over
// it boxes both operands and walks the fields at run time.
//
// Examples:
//
//	Changed(Config{Retries: 1}, Config{Retries: 2}) => true
func Changed(a, b Config) bool {
	// CHANGE CODE BELOW THIS LINE
	return !reflect.DeepEqual(a, b)
	// CHANGE CODE ABOVE THIS LINE
}
