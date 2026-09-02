// Package tablecaseptrbug — Gopher Workplace challenge.
package tablecaseptrbug

// Case is one row of a generic table test.
type Case[T any] struct {
	Name string
	In   T
}

// Pointers returns a pointer to each case, in order.
// Each pointer must address a distinct case.
//
// Examples:
//
//	Pointers(cases)[0].Name => the first case's name
func Pointers[T any](cases []Case[T]) []*Case[T] {
	// CHANGE CODE BELOW THIS LINE
	out := make([]*Case[T], 0, len(cases))
	var c Case[T]
	for _, c = range cases {
		out = append(out, &c)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
