// Package ordermerge — Gopher Workplace challenge.
package ordermerge

// MergeOrderFeeds fans every regional order feed into one merged channel and
// returns it. One goroutine per feed forwards ids as they arrive; the merged
// channel is closed exactly once, after every feed has been drained, so the
// checkout reconciler can simply range over the result.
//
// Interleaving between feeds is not defined — only that every id arrives
// exactly once and that the merged channel is closed at the end.
//
// Examples:
//
//	MergeOrderFeeds(chan eu-1,eu-2 | chan us-1) => channel yielding 3 ids, then closed
//	MergeOrderFeeds(closed empty | chan ap-1)   => channel yielding ap-1, then closed
//	MergeOrderFeeds()                           => already closed, empty channel
func MergeOrderFeeds(feeds ...<-chan string) <-chan string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
