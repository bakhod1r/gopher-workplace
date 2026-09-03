// Package pipelinestage — Gopher Workplace challenge.
package pipelinestage

// Stage runs f over every value arriving on in, sending the results on the
// returned channel, and closes that channel when in closes. Each stage is one
// goroutine, so a pipeline of k stages processes k values at once.
//
// Examples:
//
//	out := Stage(in, double)
func Stage(in <-chan int, f func(int) int) <-chan int {
	panic("not implemented")
}

// Run pushes the values through every stage in order and collects the
// results. Order is preserved because a single channel between stages is
// first-in-first-out, and no goroutine may be left running when Run returns.
//
// Examples:
//
//	Run([]int{1, 2}, []func(int) int{double, inc}) => []int{3, 5}
func Run(values []int, stages []func(int) int) []int {
	panic("not implemented")
}
