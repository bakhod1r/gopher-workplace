// Package scrapecollect — Gopher Workplace challenge.
package scrapecollect

import "time"

// Sample is one metric value scraped from one target.
type Sample struct {
	Target string
	Value  float64
}

// CollectScrapes gathers up to want samples from the scrape channel, giving up
// when the budget expires or when the scrape pool closes the channel early.
// The bool reports whether the full set arrived; on a short read the samples
// gathered so far are still returned so the scrape can be recorded as partial.
//
// The timeout is a budget for the whole collection, not per sample.
//
// Examples:
//
//	CollectScrapes(chan 3 samples, 3, 5s) => 3 samples, true
//	CollectScrapes(chan 1 sample closed, 3, 5s) => 1 sample, false
//	CollectScrapes(silent chan, 2, 20ms) => no samples, false
func CollectScrapes(scrapes <-chan Sample, want int, budget time.Duration) ([]Sample, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
