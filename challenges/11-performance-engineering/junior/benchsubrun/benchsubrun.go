// Package benchsubrun — Gopher Workplace challenge.
package benchsubrun

// Names returns the full benchmark name of each sub-benchmark that
// b.Run(fmt.Sprintf("size=%d", size), ...) would create under base:
// "<base>/size=<size>", one per size, in the order given.
// An empty sizes slice yields an empty, non-nil result.
//
// Examples:
//
//	Names("BenchmarkEncode", []int{1, 10}) => ["BenchmarkEncode/size=1" "BenchmarkEncode/size=10"]
func Names(base string, sizes []int) []string {
	panic("not implemented")
}
