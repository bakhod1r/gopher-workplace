// Package asserthelperpassbug — Gopher Workplace challenge.
package asserthelperpassbug

// Reporter is the subset of *testing.T that AssertEqual needs.
type Reporter interface {
	Errorf(format string, args ...any)
}

// AssertEqual reports a mismatch on t and returns whether got equals want.
// Callers gate follow-up checks on the returned value.
//
// Examples:
//
//	if !AssertEqual(t, got, want) { return }
func AssertEqual[T comparable](t Reporter, got, want T) bool {
	// CHANGE CODE BELOW THIS LINE
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	return true
	// CHANGE CODE ABOVE THIS LINE
}
