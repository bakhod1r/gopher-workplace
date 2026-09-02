// Package filemanifest — Gopher Workplace challenge.
package filemanifest

// FileSizes stats every path concurrently, one goroutine per path, and
// returns the sizes in the same order as paths.
//
// Examples:
//
//	FileSizes([]string{"a", "bb"}, size)  => []int{1, 2}
//	FileSizes([]string{"abc"}, size)      => []int{3}
//	FileSizes(nil, size)                  => []int{}
func FileSizes(paths []string, size func(string) int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
