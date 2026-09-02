// Package thumbnailfanout — Gopher Workplace challenge.
package thumbnailfanout

// RenderThumbnails renders every image with the given number of worker
// goroutines pulling from one job channel, and returns the thumbnail keys
// sorted ascending. workers is >= 1.
//
// Examples:
//
//	RenderThumbnails([]string{"a", "b"}, 2, suffix)  => []string{"a.thumb", "b.thumb"}
//	RenderThumbnails([]string{"z"}, 4, suffix)       => []string{"z.thumb"}
//	RenderThumbnails(nil, 2, suffix)                 => nil
func RenderThumbnails(images []string, workers int, render func(string) string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
