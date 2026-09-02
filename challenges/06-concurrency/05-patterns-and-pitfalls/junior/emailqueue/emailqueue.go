// Package emailqueue — Gopher Workplace challenge.
package emailqueue

// SendCampaign delivers a campaign to every recipient concurrently while
// never letting more than limit deliveries be in flight at once, and returns
// the total of the per-send results. limit is >= 1.
//
// Examples:
//
//	SendCampaign([]string{"a@x", "bb@x"}, 2, byteCost)  => 3
//	SendCampaign([]string{"a@x"}, 1, byteCost)          => 1
//	SendCampaign(nil, 3, byteCost)                      => 0
func SendCampaign(recipients []string, limit int, send func(string) int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
