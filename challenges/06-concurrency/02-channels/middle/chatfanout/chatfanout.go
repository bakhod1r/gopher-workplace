// Package chatfanout — Gopher Workplace challenge.
package chatfanout

// FanOutRoom broadcasts every room message to each of the subscribers and
// returns one receive-only channel per subscriber. Every subscriber sees the
// full message history in publication order, and each subscriber channel is
// closed once the room feed closes so a client goroutine can range over it.
//
// The returned channels are unbuffered: a slow subscriber applies backpressure
// to the whole room, which is the intended fairness rule here.
//
// Examples:
//
//	FanOutRoom(chan hi,bye, 2) => 2 channels, each yielding hi then bye, then closed
//	FanOutRoom(chan hi, 1)     => 1 channel yielding hi, then closed
//	FanOutRoom(chan hi, 0)     => no channels; the feed is still drained
func FanOutRoom(messages <-chan string, subscribers int) []<-chan string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
