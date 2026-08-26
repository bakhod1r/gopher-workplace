package actorrouter

import "testing"

func TestRouter(t *testing.T) {
	r := &Router{
		workers: []*Worker{
			{Inbox: make(chan int, 10)},
			{Inbox: make(chan int, 10)},
		},
	}

	r.Route(1)
	r.Route(2)
	r.Route(3)

	if <-r.workers[0].Inbox != 1 {
		t.Error("worker 0 missed 1")
	}
	if <-r.workers[1].Inbox != 2 {
		t.Error("worker 1 missed 2")
	}
	if <-r.workers[0].Inbox != 3 {
		t.Error("worker 0 missed 3")
	}
}
