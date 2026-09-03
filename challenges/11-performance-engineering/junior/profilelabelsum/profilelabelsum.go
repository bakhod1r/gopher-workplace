// Package profilelabelsum — Gopher Workplace challenge.
package profilelabelsum

// Sample is one profile sample carrying pprof labels.
type Sample struct {
	Labels map[string]string
	Value  int64
}

// SumByLabel groups a profile by the value of one label key, the way
// `pprof -tagfocus` slices a profile by request kind or tenant. Samples that
// lack the key are grouped under "", and samples with a non-positive value
// are ignored.
//
// Examples:
//
//	SumByLabel([{{"kind":"read"},3},{{"kind":"read"},2}], "kind") => {"read":5}
func SumByLabel(samples []Sample, key string) map[string]int64 {
	panic("not implemented")
}
