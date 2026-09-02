// Package billingtotals - Gopher Workplace challenge.
package billingtotals

import "sync"

// Totals accumulates invoice amounts during a billing run.
type Totals struct {
	mu    sync.Mutex
	sum   int64
	count int64
}

// Add records one invoice amount.
//
// Examples:
//
//	var t Totals; t.Add(300); t.Total() => 300
func (t *Totals) Add(amount int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Total returns the summed invoice amount.
func (t *Totals) Total() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Count returns how many invoices were added.
func (t *Totals) Count() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Average returns the truncated mean amount, or 0 when nothing was added.
//
// Examples:
//
//	var t Totals; t.Add(300); t.Add(100); t.Average() => 200
//	var t Totals; t.Average()                         => 0
func (t *Totals) Average() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
