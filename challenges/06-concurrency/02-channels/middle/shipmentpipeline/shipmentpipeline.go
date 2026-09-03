// Package shipmentpipeline — Gopher Workplace challenge.
package shipmentpipeline

// Labels runs orders through a two-stage pipeline: keep first, then render.
// Each stage runs in its own goroutine and passes work on through a channel;
// orders rejected by keep never reach render. Results come back in input
// order.
//
// Examples:
//
//	Labels([]string{"o1","o2"}, all, upper)  => ["LABEL-o1" "LABEL-o2"]
//	Labels([]string{"o1","bad"}, notBad, up) => ["LABEL-o1"]
//	Labels(nil, all, upper)                  => []
func Labels(orders []string, keep func(order string) bool, render func(order string) string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
