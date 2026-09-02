// Package thumbnailer — Gopher Workplace challenge.
package thumbnailer

// Image is a source image with its pixel dimensions.
type Image struct {
	Width  int
	Height int
}

// TargetHeights returns the proportional height of every image scaled to maxWidth.
//
// Examples:
//
//	TargetHeights([]Image{{100, 50}, {200, 200}}, 100)  => [50 100]
//	TargetHeights([]Image{{0, 40}}, 100)                => [0]
//	TargetHeights(nil, 100)                             => []
func TargetHeights(images []Image, maxWidth int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
