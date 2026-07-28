// Package mapptrupdate stores accounts in a map and credits one. A planted bug
// fetches a COPY of the pointer's pointee, mutates the copy, and stores nothing,
// so the update is lost. Mutate through the fetched pointer instead.
package mapptrupdate

type Account struct{ Balance int }

// Credit adds amt to the account stored at id. Returns false if id is absent.
func Credit(m map[int]*Account, id, amt int) bool {
	a, ok := m[id]
	if !ok {
		return false
	}
	// CHANGE CODE BELOW THIS LINE
	acc := *a
	acc.Balance += amt
	// CHANGE CODE ABOVE THIS LINE
	return true
}
