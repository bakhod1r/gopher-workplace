package eventbusgen

import "testing"

func TestBusPublish(t *testing.T) {
	var b Bus[int]
	got := 0
	b.Subscribe(func(v int) { got += v })
	if n := b.Publish(5); n != 1 {
		t.Errorf("Publish returned %d, want 1", n)
	}
	if got != 5 {
		t.Errorf("handler saw %d, want 5", got)
	}
}

func TestBusUnsubscribe(t *testing.T) {
	var b Bus[string]
	calls := 0
	id := b.Subscribe(func(string) { calls++ })
	b.Publish("x")
	if !b.Unsubscribe(id) {
		t.Fatal("Unsubscribe = false, want true")
	}
	if n := b.Publish("y"); n != 0 {
		t.Errorf("Publish after unsubscribe returned %d, want 0", n)
	}
	if calls != 1 {
		t.Errorf("handler called %d times, want 1", calls)
	}
	if b.Unsubscribe(id) {
		t.Error("Unsubscribe on an unknown id = true, want false")
	}
}

func TestBusIdsAreUnique(t *testing.T) {
	var b Bus[int]
	a := b.Subscribe(func(int) {})
	c := b.Subscribe(func(int) {})
	if a == c {
		t.Error("two subscriptions share an id")
	}
	b.Unsubscribe(a)
	d := b.Subscribe(func(int) {})
	if d == a {
		t.Error("an id was reused after unsubscribing")
	}
	if n := b.Publish(1); n != 2 {
		t.Errorf("Publish returned %d, want 2", n)
	}
}
