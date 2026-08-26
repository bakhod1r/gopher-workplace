package pubsub

import "testing"

func TestPubSub(t *testing.T) {
	ps := New()
	ch1 := ps.Subscribe("news")
	ch2 := ps.Subscribe("news")

	ps.Publish("news", "hello")

	if got := <-ch1; got != "hello" {
		t.Errorf("ch1 got %q", got)
	}
	if got := <-ch2; got != "hello" {
		t.Errorf("ch2 got %q", got)
	}
}
