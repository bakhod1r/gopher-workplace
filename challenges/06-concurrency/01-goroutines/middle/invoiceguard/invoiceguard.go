// Package invoiceguard — Gopher Workplace challenge.
package invoiceguard

// Rendered is the outcome of rendering one invoice.
type Rendered struct {
	ID  string
	Doc string
	Err error
}

// RenderInvoices renders every invoice in its own goroutine and returns one
// Rendered per ID, in input order. A template that panics on one malformed
// invoice must not take the whole billing run down with it: the worker recovers
// the panic and reports it as that invoice's Err.
//
// Examples:
//
//	RenderInvoices([]string{"INV-1"}, render)         => [{INV-1 doc:INV-1 <nil>}]
//	RenderInvoices([]string{"INV-BAD"}, render)       => [{INV-BAD  panic}]
//	RenderInvoices(nil, render)                       => []
func RenderInvoices(ids []string, render func(id string) string) []Rendered {
	// TODO(candidate): implement this.
	panic("not implemented")
}
