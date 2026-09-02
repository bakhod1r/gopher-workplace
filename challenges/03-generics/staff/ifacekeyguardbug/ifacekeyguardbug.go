// Package ifacekeyguardbug — Gopher Workplace challenge.
package ifacekeyguardbug

// Distinct returns the elements of vals in input order, dropping later duplicates.
// An element whose dynamic type cannot be used as a map key is never deduplicated;
// it is passed through unchanged instead of crashing the caller.
//
// Examples:
//
//	Distinct([]any{1, 1, []int{2}}) => []any{1, []int{2}}
func Distinct[T comparable](vals []T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(vals))
	seen := make(map[T]struct{}, len(vals))
	for _, v := range vals {
		if _, dup := seen[v]; dup {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}

// hashable reports whether v can be used as a map key at run time.
func hashable[T comparable](v T) (ok bool) {
	defer func() {
		if recover() != nil {
			ok = false
		}
	}()
	probe := make(map[any]struct{}, 1)
	probe[any(v)] = struct{}{}
	return true
}
