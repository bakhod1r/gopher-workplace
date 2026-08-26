package chainptr

import "testing"

func TestChain(t *testing.T) {
	c := NewConfig().Set("host", "localhost").Set("port", "8080")

	cases := []struct {
		key  string
		want string
	}{
		{"host", "localhost"},
		{"port", "8080"},
	}

	for _, tc := range cases {
		t.Run(tc.key, func(t *testing.T) {
			if got := c.Data[tc.key]; got != tc.want {
				t.Errorf("Data[%q] = %q, want %q", tc.key, got, tc.want)
			}
		})
	}

	t.Run("overwrite", func(t *testing.T) {
		c.Set("host", "example.com")
		if got := c.Data["host"]; got != "example.com" {
			t.Errorf("after overwrite: Data[host] = %q, want example.com", got)
		}
	})
}
