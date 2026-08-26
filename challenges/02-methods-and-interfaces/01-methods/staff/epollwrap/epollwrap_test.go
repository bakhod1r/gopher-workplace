package epollwrap

import "testing"

func TestEpoll(t *testing.T) {
	e := &Epoll{Active: true}
	if !e.Wait() {
		t.Error("expected true")
	}
}
