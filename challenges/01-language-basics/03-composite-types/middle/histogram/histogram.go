// Package histogram buckets values into fixed-width bins.
package histogram

// Bucket returns counts per bin of width size for values in [0, +inf). Bin i
// covers [i*size, (i+1)*size). The result length is enough to hold the max
// value; empty input returns an empty slice. size must be > 0.
//
// TODO(candidate): find max bin, allocate, count.
func Bucket(xs []int, size int) []int {
	panic("not implemented")
}
