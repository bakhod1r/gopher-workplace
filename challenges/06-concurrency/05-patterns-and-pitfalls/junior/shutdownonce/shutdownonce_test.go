package shutdownonce

import "testing"

func TestShutdownOnce(t *testing.T) {
	cases := []struct {
		name       string
		closers    int
		wantClosed bool
	}{
		{"one_closer", 1, true},
		{"two_closers", 2, true},
		{"ten_closers", 10, true},
		{"fifty_closers", 50, true},
		{"no_closers", 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			quit := make(chan struct{})
			ShutdownOnce(quit, tc.closers)

			closed := false
			select {
			case <-quit:
				closed = true
			default:
			}

			if closed != tc.wantClosed {
				t.Errorf("ShutdownOnce(quit, %d): closed = %v, want %v", tc.closers, closed, tc.wantClosed)
			}
		})
	}
}
