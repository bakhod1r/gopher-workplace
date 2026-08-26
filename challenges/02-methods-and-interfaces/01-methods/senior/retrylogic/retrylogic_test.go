package retrylogic

import "testing"

func TestRetry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := &Client{FailInt: 2}
		err := c.DoWithRetry(3)
		if err != nil {
			t.Errorf("expected success, got %v", err)
		}
		if c.Attempts != 3 {
			t.Errorf("Attempts = %d, want 3", c.Attempts)
		}
	})

	t.Run("exhausted", func(t *testing.T) {
		c := &Client{FailInt: 5}
		err := c.DoWithRetry(3)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if c.Attempts != 3 {
			t.Errorf("Attempts = %d, want 3", c.Attempts)
		}
	})
}
