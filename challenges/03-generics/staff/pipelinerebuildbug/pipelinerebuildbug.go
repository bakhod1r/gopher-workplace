// Package pipelinerebuildbug — Gopher Workplace challenge.
package pipelinerebuildbug

// Pipeline applies each stage to every element, in stage order.
// It allocates one buffer regardless of how many stages there are.
//
// Examples:
//
//	Pipeline([]int{1, 2}, inc, double) => []int{4, 6}
func Pipeline[T any](s []T, stages ...func(T) T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, len(s))
	copy(out, s)
	for _, st := range stages {
		next := make([]T, 0, len(out))
		for _, v := range out {
			next = append(next, st(v))
		}
		out = next
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
