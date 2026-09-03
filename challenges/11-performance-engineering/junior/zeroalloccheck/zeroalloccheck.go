// Package zeroalloccheck — Gopher Workplace challenge.
package zeroalloccheck

// Result reports what an allocation check found.
type Result struct {
	Allocs int
	Limit  int
	OK     bool
}

// Check runs f repeatedly, measures its average allocations per call, and
// reports whether that stays within limit. The measured count is rounded to
// the nearest whole allocation, a negative limit is treated as 0, and runs is
// clamped to at least 1.
//
// Examples:
//
//	Check(100, 0, func() {}) => Result{Allocs: 0, Limit: 0, OK: true}
func Check(runs, limit int, f func()) Result {
	panic("not implemented")
}
