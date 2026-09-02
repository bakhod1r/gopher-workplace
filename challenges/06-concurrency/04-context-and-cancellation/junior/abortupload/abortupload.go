// Package abortupload — Gopher Workplace challenge.
package abortupload

// AbortUpload starts a cancellable upload context, aborts it because the user
// pressed Cancel in the UI, and returns the reason recorded on the context.
//
// Examples:
//
//	AbortUpload()                                     => context.Canceled
//	errors.Is(AbortUpload(), context.Canceled)        => true
//	errors.Is(AbortUpload(), context.DeadlineExceeded) => false
func AbortUpload() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
