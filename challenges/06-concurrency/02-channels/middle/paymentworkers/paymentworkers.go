// Package paymentworkers — Gopher Workplace challenge.
package paymentworkers

// Result is one capture attempt's outcome.
type Result struct {
	Charge string
	Status string
}

// CaptureAll runs workers goroutines that consume charges and call capture on
// each one, returning every outcome keyed by charge ID. A workers count of
// zero or less means one worker.
//
// Examples:
//
//	CaptureAll([]string{"ch_1"}, 2, ok)          => map[ch_1:captured]
//	CaptureAll([]string{"ch_1","ch_2"}, 1, ok)   => map[ch_1:captured ch_2:captured]
//	CaptureAll(nil, 4, ok)                       => map[]
func CaptureAll(charges []string, workers int, capture func(charge string) string) map[string]string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
