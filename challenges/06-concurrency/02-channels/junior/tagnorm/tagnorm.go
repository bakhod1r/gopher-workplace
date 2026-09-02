// Package tagnorm — Gopher Workplace challenge.
package tagnorm

// NormalizeTags receives every label value from in, sends the upper-cased
// value on out, and closes out when in is exhausted.
//
// Examples:
//
//	NormalizeTags(in "az","eu", out) => out yields "AZ", "EU" then closed
//	NormalizeTags(in "Prod", out)    => out yields "PROD" then closed
//	NormalizeTags(in closed, out)    => out closed immediately
func NormalizeTags(in <-chan string, out chan<- string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
