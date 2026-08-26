package asyncfetch

import (
	"testing"
	"time"
)

func TestFetchAsync(t *testing.T) {
	f := &Fetcher{}
	done := f.FetchAsync("abc")

	select {
	case <-done:
		if f.Result != "data: abc" {
			t.Errorf("Result = %q, want %q", f.Result, "data: abc")
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for FetchAsync")
	}
}
