// Package thumbnailworkers — Gopher Workplace challenge.
package thumbnailworkers

import "context"

// Upload is one image waiting in the upload queue.
type Upload struct {
	Key string
}

// Thumbnail is the rendered result for an upload.
type Thumbnail struct {
	Key   string
	Width int
}

// RenderQueue drains the upload queue with a pool of worker goroutines and
// returns the thumbnails in the same order as uploads. The pool shares one
// cancellable context: the first render error cancels it, so the remaining
// workers stop pulling new jobs instead of burning CPU on a batch that is
// already doomed.
//
// It returns ctx.Err() if the caller's context is already finished, the first
// render error if any upload failed, or the ordered thumbnails.
//
// Examples:
//
//	RenderQueue(live ctx, 3 uploads, 2 workers, render)   => 3 thumbnails in order
//	RenderQueue(live ctx, uploads with "bad", 2, render)  => errRender
//	RenderQueue(cancelled ctx, 3 uploads, 2, render)      => context.Canceled
func RenderQueue(ctx context.Context, uploads []Upload, workers int, render func(context.Context, Upload) (Thumbnail, error)) ([]Thumbnail, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
