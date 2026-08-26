// Package stmtransact — Gopher Workplace challenge.
package stmtransact

// TVar is a transactional variable.
type TVar struct {
	val int
}

// Tx is a transaction block.
func Tx(tv *TVar, fn func(int) int) {
	// TODO(candidate): read tv.val, pass to fn, assign result to tv.val
	// No locks needed for this basic simulation.
	panic("not implemented")
}
