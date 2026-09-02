package auditsplit

import (
	"sync"
	"testing"
)

func TestTeeAudit(t *testing.T) {
	send := func(vals ...string) <-chan string {
		ch := make(chan string, len(vals))
		for _, v := range vals {
			ch <- v
		}
		close(ch)
		return ch
	}

	cases := []struct {
		name   string
		events []string
	}{
		{"single_event", []string{"login"}},
		{"three_events", []string{"login", "update", "logout"}},
		{"duplicate_events", []string{"login", "login"}},
		{"five_events", []string{"a", "b", "c", "d", "e"}},
		{"empty_stream", nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			archive, alerts := TeeAudit(send(tc.events...))

			var (
				wg       sync.WaitGroup
				gotArch  []string
				gotAlert []string
			)
			wg.Add(2)
			go func() {
				defer wg.Done()
				for ev := range archive {
					gotArch = append(gotArch, ev)
				}
			}()
			go func() {
				defer wg.Done()
				for ev := range alerts {
					gotAlert = append(gotAlert, ev)
				}
			}()
			wg.Wait()

			for _, got := range [][]string{gotArch, gotAlert} {
				if len(got) != len(tc.events) {
					t.Fatalf("TeeAudit(%v) delivered %v, want %v", tc.events, got, tc.events)
				}
				for i := range got {
					if got[i] != tc.events[i] {
						t.Fatalf("TeeAudit(%v) delivered %v, want %v", tc.events, got, tc.events)
					}
				}
			}
		})
	}
}
