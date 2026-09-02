// Package uploadfeed — Gopher Workplace challenge.
package uploadfeed

// UploadFeed streams the object keys of a batch upload on a fresh channel,
// closes it when the batch is exhausted, and returns the receive-only end.
//
// Examples:
//
//	UploadFeed([]string{"a.jpg", "b.jpg"})  => yields "a.jpg", "b.jpg" then closes
//	UploadFeed([]string{"only.png"})        => yields "only.png" then closes
//	UploadFeed(nil)                         => closes immediately
func UploadFeed(keys []string) <-chan string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
