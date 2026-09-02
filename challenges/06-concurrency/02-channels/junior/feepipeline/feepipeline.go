// Package feepipeline — Gopher Workplace challenge.
package feepipeline

// LineTotals pushes unit counts through two goroutine stages — two cents a
// unit, then a one-cent handling fee — and returns the line totals in order.
//
// Each unit count n becomes n*2+1 cents.
//
// Examples:
//
//	LineTotals([]int{1,2}) => [3 5]
//	LineTotals([]int{0}) => [1]
//	LineTotals(nil) => []
func LineTotals(units []int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
