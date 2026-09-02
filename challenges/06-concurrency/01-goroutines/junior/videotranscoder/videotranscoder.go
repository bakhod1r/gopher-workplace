// Package videotranscoder — Gopher Workplace challenge.
package videotranscoder

// TargetBitrates scales every source bitrate down to the requested percentage.
//
// Examples:
//
//	TargetBitrates([]int{4000, 2000}, 50)  => [2000 1000]
//	TargetBitrates([]int{4000}, 100)       => [4000]
//	TargetBitrates(nil, 50)                => []
func TargetBitrates(bitrates []int, factorPct int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
