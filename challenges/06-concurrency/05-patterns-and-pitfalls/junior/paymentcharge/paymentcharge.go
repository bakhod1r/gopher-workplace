// Package paymentcharge — Gopher Workplace challenge.
package paymentcharge

// ChargeAll charges every payment provider concurrently and returns the
// messages of the failures, sorted ascending. Providers that succeed
// contribute nothing.
//
// Examples:
//
//	ChargeAll([]string{"ok-visa", "bad-amex"}, charge)  => []string{"bad-amex declined"}
//	ChargeAll([]string{"ok-visa"}, charge)              => nil
//	ChargeAll(nil, charge)                              => nil
func ChargeAll(providers []string, charge func(string) error) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
