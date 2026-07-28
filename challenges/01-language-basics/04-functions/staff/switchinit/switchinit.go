// Package switchinit classifies the sign of an expensive computation, counting
// how many times the computation runs. A planted bug recomputes it in each case
// instead of using a switch init statement, so the count is wrong (and the logic
// duplicated). Use `switch v := f(); { ... }` to compute once.
package switchinit

// classify returns the sign label of compute(x) and how many times compute ran.
var calls int

func compute(x int) int { calls++; return x }

// Classify returns the sign of compute(x) computing it EXACTLY once.
func Classify(x int) (label string, runs int) {
	calls = 0
	// CHANGE CODE BELOW THIS LINE
	switch {
	case compute(x) < 0:
		label = "neg"
	case compute(x) == 0:
		label = "zero"
	default:
		label = "pos"
	}
	// CHANGE CODE ABOVE THIS LINE
	return label, calls
}
