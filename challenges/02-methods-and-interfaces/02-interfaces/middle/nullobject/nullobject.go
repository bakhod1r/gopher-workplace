// Package nullobject — Gopher Workplace challenge.
package nullobject

// Metrics records named events.
type Metrics interface {
	Report(name string)
}

// Recorder keeps every reported metric.
type Recorder struct {
	Events []string
}

// Report records the metric.
func (r *Recorder) Report(name string) {
	// TODO(candidate): append the event.
	panic("not implemented")
}

// NopMetrics discards every metric.
type NopMetrics struct{}

// Report does nothing.
func (NopMetrics) Report(name string) {
	// TODO(candidate): do nothing.
	panic("not implemented")
}

// MetricsOr returns m, or NopMetrics{} when m is nil.
//
// Examples:
//
//	MetricsOr(nil) => NopMetrics{}
func MetricsOr(m Metrics) Metrics {
	// TODO(candidate): normalise the nil case.
	panic("not implemented")
}

// Process reports "item:<item>" for each item and returns the count.
// Its body must not contain a nil check.
func Process(m Metrics, items []string) int {
	// TODO(candidate): normalise once, then report freely.
	panic("not implemented")
}
