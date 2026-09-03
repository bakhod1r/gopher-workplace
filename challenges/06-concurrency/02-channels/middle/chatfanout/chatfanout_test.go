package chatfanout

import (
	"reflect"
	"sync"
	"testing"
)

func roomOf(msgs ...string) <-chan string {
	ch := make(chan string, len(msgs))
	for _, m := range msgs {
		ch <- m
	}
	close(ch)
	return ch
}

func TestFanOutRoom(t *testing.T) {
	cases := []struct {
		name        string
		messages    []string
		subscribers int
	}{
		{"two_subscribers", []string{"hi", "bye"}, 2},
		{"one_subscriber", []string{"hi"}, 1},
		{"no_subscribers", []string{"hi", "bye"}, 0},
		{"empty_room", nil, 3},
		{"four_subscribers", []string{"a", "b", "c"}, 4},
		{"single_message_many_subs", []string{"deploy"}, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			outs := FanOutRoom(roomOf(tc.messages...), tc.subscribers)
			if len(outs) != tc.subscribers {
				t.Fatalf("got %d subscriber channels, want %d", len(outs), tc.subscribers)
			}

			want := tc.messages
			if want == nil {
				want = []string{}
			}

			got := make([][]string, tc.subscribers)
			var wg sync.WaitGroup
			wg.Add(tc.subscribers)
			for i, out := range outs {
				go func(i int, out <-chan string) {
					defer wg.Done()
					seen := []string{}
					for m := range out {
						seen = append(seen, m)
					}
					got[i] = seen
				}(i, out)
			}
			wg.Wait()

			for i, seen := range got {
				if !reflect.DeepEqual(seen, want) {
					t.Errorf("subscriber %d saw %#v, want %#v", i, seen, want)
				}
			}
		})
	}
}

func TestFanOutRoomDrainsWithoutSubscribers(t *testing.T) {
	room := make(chan string, 2)
	room <- "hi"
	room <- "bye"
	close(room)

	if outs := FanOutRoom(room, 0); len(outs) != 0 {
		t.Fatalf("got %d channels, want 0", len(outs))
	}
}
