// Package spreadmissing sums a slice by delegating to a variadic helper. A
// planted bug builds a nested slice instead of spreading, so it sums one value.
package spreadmissing

func sum(nums ...int) int {
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}

// Total sums all elements of xs via the variadic sum helper.
func Total(xs []int) int {
	// CHANGE CODE BELOW THIS LINE
	return sum(len(xs))
	// CHANGE CODE ABOVE THIS LINE
}
