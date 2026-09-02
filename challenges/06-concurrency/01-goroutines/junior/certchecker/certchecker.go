// Package certchecker — Gopher Workplace challenge.
package certchecker

// ExpiringSoon flags every certificate that expires within the alert window.
//
// Examples:
//
//	ExpiringSoon([]int{100, 400}, 90, 30)  => [true false]
//	ExpiringSoon([]int{50}, 90, 30)        => [true]
//	ExpiringSoon(nil, 90, 30)              => []
func ExpiringSoon(expiries []int, today int, window int) []bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
