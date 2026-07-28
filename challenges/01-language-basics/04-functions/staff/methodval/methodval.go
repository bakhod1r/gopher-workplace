// Package methodval captures a bound method value. A planted bug rebinds the
// method to a later receiver state, but a method VALUE snapshots the receiver
// when created — the test pins that a value bound early sees the early state.
package methodval

type counter struct{ n int }

func (c counter) Get() int { return c.n }

// BoundEarly returns a function that reports the counter's value AS IT WAS when
// the function was created (n = start), even though the local is later changed.
func BoundEarly(start int) func() int {
	c := counter{n: start}
	f := c.Get // method value: receiver copied now
	c.n = 999  // must NOT affect f (value receiver was copied)
	// CHANGE CODE BELOW THIS LINE
	f = c.Get
	// CHANGE CODE ABOVE THIS LINE
	return f
}
