// Package stdsortstableweakbug — Gopher Workplace challenge.
package stdsortstableweakbug

import (
	"slices"
)

// Task is one unit of queued work.
type Task struct {
	Name string
	Pri  int
}

// SortByPriority returns tasks ordered by ascending Pri.
// Tasks with equal Pri keep their input order. The input is not modified.
//
// Examples:
//
//	SortByPriority([{a 2} {b 1} {c 2}]) => [{b 1} {a 2} {c 2}]
func SortByPriority(tasks []Task) []Task {
	// CHANGE CODE BELOW THIS LINE
	out := slices.Clone(tasks)
	slices.SortStableFunc(out, func(a, b Task) int {
		if a.Pri > b.Pri {
			return 1
		}
		return -1
	})
	return out
	// CHANGE CODE ABOVE THIS LINE
}
