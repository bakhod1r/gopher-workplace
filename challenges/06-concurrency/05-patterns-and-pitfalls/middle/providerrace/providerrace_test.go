package providerrace

import (
	"context"
	"errors"
	"strings"
	"testing"
)

var errDeclined = errors.New("declined")

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

// authorize succeeds for "ok-" providers, declines "no-" providers, and blocks
// until the context is cancelled for "slow-" providers.
func authorize(ctx context.Context, provider string, c Charge) (string, error) {
	switch {
	case strings.HasPrefix(provider, "ok-"):
		return provider + ":auth", nil
	case strings.HasPrefix(provider, "slow-"):
		<-ctx.Done()
		return "", ctx.Err()
	default:
		return "", errDeclined
	}
}

func TestFirstAuthorization(t *testing.T) {
	charge := Charge{Amount: 1250, Currency: "EUR"}
	live := context.Background()

	cases := []struct {
		name      string
		ctx       context.Context
		providers []string
		want      string
		wantErr   error
	}{
		{"winner_between_slow_providers", live, []string{"slow-a", "ok-b", "slow-c"}, "ok-b:auth", nil},
		{"winner_after_declines", live, []string{"no-a", "ok-b", "no-c"}, "ok-b:auth", nil},
		{"single_provider", live, []string{"ok-only"}, "ok-only:auth", nil},
		{"all_decline", live, []string{"no-a", "no-b"}, "", ErrAllProvidersDeclined},
		{"no_providers", live, nil, "", ErrAllProvidersDeclined},
		{"checkout_abandoned", cancelled(), []string{"ok-a"}, "", context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FirstAuthorization(tc.ctx, charge, tc.providers, authorize)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("FirstAuthorization(%v) error = %v, want %v", tc.providers, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("FirstAuthorization(%v) = %q, want %q", tc.providers, got, tc.want)
			}
		})
	}
}

func TestFirstAuthorizationCancelsLosers(t *testing.T) {
	charge := Charge{Amount: 100, Currency: "USD"}
	// Every slow provider blocks until the shared context is cancelled. If the
	// winner does not cancel them, this test never returns.
	providers := []string{"slow-1", "slow-2", "slow-3", "ok-fast", "slow-4"}
	got, err := FirstAuthorization(context.Background(), charge, providers, authorize)
	if err != nil {
		t.Fatalf("FirstAuthorization() error = %v", err)
	}
	if got != "ok-fast:auth" {
		t.Errorf("FirstAuthorization() = %q, want %q", got, "ok-fast:auth")
	}
}
