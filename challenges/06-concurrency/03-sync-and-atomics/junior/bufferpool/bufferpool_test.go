package bufferpool

import (
	"strings"
	"sync"
	"testing"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		name   string
		fields []string
		want   string
	}{
		{"two_fields", []string{"warn", "disk full"}, "warn|disk full"},
		{"single", []string{"solo"}, "solo"},
		{"empty_slice", []string{}, ""},
		{"nil", nil, ""},
		{"three", []string{"a", "b", "c"}, "a|b|c"},
		{"empty_field", []string{"", "x"}, "|x"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := NewEncoder()
			if got := e.Encode(tc.fields); got != tc.want {
				t.Errorf("Encode(%v) = %q, want %q", tc.fields, got, tc.want)
			}
		})
	}
}

func TestEncodeReusedBufferIsReset(t *testing.T) {
	e := NewEncoder()
	e.Encode([]string{"aaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbb"})
	if got := e.Encode([]string{"short"}); got != "short" {
		t.Errorf("reused buffer leaked earlier data: got %q, want %q", got, "short")
	}
}

func TestEncodeConcurrent(t *testing.T) {
	e := NewEncoder()
	const n = 200

	var wg sync.WaitGroup
	got := make([]string, n)
	for i := range n {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			got[i] = e.Encode([]string{"line", strings.Repeat("x", i%7)})
		}(i)
	}
	wg.Wait()

	for i, s := range got {
		want := "line|" + strings.Repeat("x", i%7)
		if s != want {
			t.Fatalf("goroutine %d: got %q, want %q", i, s, want)
		}
	}
}
