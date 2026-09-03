// Package stringkeycollidebug — Gopher Workplace challenge.
package stringkeycollidebug

import "fmt"

// Distinct returns the elements of vals in input order, dropping later duplicates.
//
// Two elements are duplicates only when they are equal under ==. Elements that
// merely share a textual rendering are distinct and must both survive.
//
// Examples:
//
//	Distinct([]any{1, 1, 2})     => []any{1, 2}
//	Distinct([]any{1, "1"})      => []any{1, "1"}
//	Distinct([]string{"a", "a"}) => []string{"a"}
func Distinct[T comparable](vals []T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(vals))
	seen := make(map[string]struct{}, len(vals))
	for _, v := range vals {
		k := fmt.Sprint(v)
		if _, dup := seen[k]; dup {
			continue
		}
		seen[k] = struct{}{}
		out = append(out, v)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
