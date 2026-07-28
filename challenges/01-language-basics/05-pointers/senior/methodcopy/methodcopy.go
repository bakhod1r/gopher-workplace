// Package methodcopy returns a getter bound to an account. A planted bug binds
// the method value on a COPY (dereferenced value), so later deposits are not
// seen. Binding on the pointer captures the live object.
package methodcopy

type Account struct{ Balance int }

func (a *Account) Get() int { return a.Balance }

// Getter returns a function that reports a's CURRENT balance, reflecting later
// deposits.
func Getter(a *Account) func() int {
	// CHANGE CODE BELOW THIS LINE
	tmp := *a
	return tmp.Get
	// CHANGE CODE ABOVE THIS LINE
}
