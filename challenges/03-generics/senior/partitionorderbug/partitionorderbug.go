// Package partitionorderbug — Gopher Workplace challenge.
package partitionorderbug

// Partition splits s into elements that satisfy pred and the rest.
// Both halves keep their input order.
//
// Examples:
//
//	Partition([]int{1, 2, 3}, even) => [2], [1 3]
func Partition[T any](s []T, pred func(T) bool) (yes, no []T) {
	// CHANGE CODE BELOW THIS LINE
	yes, no = make([]T, 0), make([]T, 0)
	for _, v := range s {
		if pred(v) {
			yes = append(yes, v)
		} else {
			no = append([]T{v}, no...)
		}
	}
	return yes, no
	// CHANGE CODE ABOVE THIS LINE
}
