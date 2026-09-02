// Package anyboxhotloopbug — Gopher Workplace challenge.
package anyboxhotloopbug

// Number is any type Sum can add.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum adds every element of s.
// It performs no per-element allocation.
//
// Examples:
//
//	Sum([]int{1, 2, 3}) => 6
func Sum[T Number](s []T) T {
	// CHANGE CODE BELOW THIS LINE
	boxed := make([]any, 0, len(s))
	for _, v := range s {
		boxed = append(boxed, v)
	}
	var total T
	for _, b := range boxed {
		total += b.(T)
	}
	return total
	// CHANGE CODE ABOVE THIS LINE
}
