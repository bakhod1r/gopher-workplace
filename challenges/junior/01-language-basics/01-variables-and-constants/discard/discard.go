// Package discard — Gopher Workplace challenge.
package discard

// Split reports how many whole pages of the given size fit in n items, and how
// many items are left over.
//
// Callers often want only one of the two. Go has no "ignore the rest" rule: a
// multi-value call must be received in full, and the blank identifier `_` is
// how you throw a value away without naming it.
//
// This one is written for you — the exercise is in the two functions below.
//
// Examples:
//
//	Split(10, 3) => 3, 1
//	Split(9, 3)  => 3, 0
//	Split(2, 5)  => 0, 2
func Split(n, size int) (pages, rest int) {
	if size <= 0 {
		return 0, n
	}
	return n / size, n % size
}

// Pages returns only the number of whole pages, discarding the remainder.
//
// TODO(candidate): call Split and keep the first value only. You cannot write
// `pages := Split(n, size)` — the call yields two values and both must be
// received. Use the blank identifier for the one you do not want.
//
// Examples:
//
//	Pages(10, 3) => 3
//	Pages(9, 3)  => 3
//	Pages(2, 5)  => 0
func Pages(n, size int) int {
	// TODO(candidate): implement this using Split.
	panic("not implemented")
}

// Leftover returns only the remainder, discarding the page count.
//
// Examples:
//
//	Leftover(10, 3) => 1
//	Leftover(9, 3)  => 0
//	Leftover(2, 5)  => 2
func Leftover(n, size int) int {
	// TODO(candidate): implement this using Split.
	panic("not implemented")
}
