// Package methodexpr builds a function from a method expression. A planted bug
// uses the wrong form, dropping the explicit receiver argument that a method
// EXPRESSION requires.
package methodexpr

type Counter struct{ N int }

func (c *Counter) Add(d int) { c.N += d }

// AdderExpr returns the method EXPRESSION Counter.Add as a function taking the
// receiver explicitly: f(c, d) adds d to c.
func AdderExpr() func(*Counter, int) {
	// CHANGE CODE BELOW THIS LINE
	var c Counter
	return func(_ *Counter, d int) { c.Add(d) }
	// CHANGE CODE ABOVE THIS LINE
}
