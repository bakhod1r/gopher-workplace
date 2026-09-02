// Package startupsequence — Gopher Workplace challenge.
package startupsequence

// RunChecks runs every preflight check and reports its status code in order.
//
// Examples:
//
//	RunChecks([]func() int{configOK, diskFull})  => [0 28]
//	RunChecks([]func() int{configOK})            => [0]
//	RunChecks(nil)                               => []
func RunChecks(checks []func() int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
