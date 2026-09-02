// Package invoicerenderer — Gopher Workplace challenge.
package invoicerenderer

// Invoice is a billing document made of line-item amounts in cents.
type Invoice struct {
	Lines []int
}

// InvoiceTotals returns the total of every invoice, in input order.
//
// Examples:
//
//	InvoiceTotals([]Invoice{{[]int{100, 250}}})  => [350]
//	InvoiceTotals([]Invoice{{nil}})              => [0]
//	InvoiceTotals(nil)                           => []
func InvoiceTotals(invoices []Invoice) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
