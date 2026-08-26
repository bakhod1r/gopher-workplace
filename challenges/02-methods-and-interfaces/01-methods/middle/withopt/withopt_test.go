package withopt

import "testing"

func TestNewServer(t *testing.T) {
	cases := []struct {
		name string
		opts []Option
		want Server
	}{
		{"defaults", nil, Server{"localhost", 8080, 30}},
		{"custom_port", []Option{WithPort(9090)}, Server{"localhost", 9090, 30}},
		{"custom_host", []Option{WithHost("0.0.0.0")}, Server{"0.0.0.0", 8080, 30}},
		{"custom_timeout", []Option{WithTimeout(60)}, Server{"localhost", 8080, 60}},
		{"all_custom", []Option{WithHost("api"), WithPort(3000), WithTimeout(10)},
			Server{"api", 3000, 10}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewServer(tc.opts...)
			if got != tc.want {
				t.Errorf("NewServer() = %+v, want %+v", got, tc.want)
			}
		})
	}
}
