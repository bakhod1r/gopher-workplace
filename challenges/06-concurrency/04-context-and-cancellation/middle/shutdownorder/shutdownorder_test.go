package shutdownorder

import (
	"context"
	"errors"
	"strings"
	"testing"
)

var errStopFailed = errors.New("stop failed")

func cleanService(name string) Service {
	return Service{Name: name, Stop: func(ctx context.Context) error { return nil }}
}

func stuckService(name string) Service {
	return Service{Name: name, Stop: func(ctx context.Context) error { return errStopFailed }}
}

func sigterm() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func drainWindowClosed() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	cancel()
	return ctx
}

func TestShutdownServices(t *testing.T) {
	full := []Service{cleanService("database"), cleanService("cache"), cleanService("http")}

	cases := []struct {
		name        string
		ctx         context.Context
		services    []Service
		wantStopped string
		wantErr     error
	}{
		{"nothing_running", context.Background(), nil, "", nil},
		{"reverse_startup_order", context.Background(), full, "http,cache,database", nil},
		{"single_service", context.Background(), []Service{cleanService("http")}, "http", nil},
		{"middle_service_stuck", context.Background(), []Service{cleanService("database"), stuckService("cache"), cleanService("http")}, "http", errStopFailed},
		{"first_stop_fails", context.Background(), []Service{cleanService("database"), cleanService("cache"), stuckService("http")}, "", errStopFailed},
		{"sigterm_before_drain", sigterm(), full, "", context.Canceled},
		{"drain_window_closed", drainWindowClosed(), full, "", context.DeadlineExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			stopped, err := ShutdownServices(tc.ctx, tc.services)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ShutdownServices() error = %v, want %v", err, tc.wantErr)
			}
			if got := strings.Join(stopped, ","); got != tc.wantStopped {
				t.Errorf("stopped = %q, want %q", got, tc.wantStopped)
			}
			if stopped == nil {
				t.Error("stopped slice is nil, want a non-nil (possibly empty) slice")
			}
		})
	}
}
