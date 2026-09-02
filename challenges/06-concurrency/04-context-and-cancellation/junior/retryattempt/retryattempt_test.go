package retryattempt

import (
	"context"
	"testing"
)

func TestAttempt(t *testing.T) {
	bg := context.Background()

	cases := []struct {
		name string
		ctx  context.Context
		want int
	}{
		{"never_set", bg, 0},
		{"first_try", WithAttempt(bg, 1), 1},
		{"explicit_zero", WithAttempt(bg, 0), 0},
		{"second_try_shadows_first", WithAttempt(WithAttempt(bg, 1), 2), 2},
		{"third_try_shadows_both", WithAttempt(WithAttempt(WithAttempt(bg, 1), 2), 3), 3},
		{"wrong_type_defaults", context.WithValue(bg, attemptKey{}, "2"), 0},
		{"survives_derivation", func() context.Context {
			ctx, cancel := context.WithCancel(WithAttempt(bg, 4))
			defer cancel()
			return ctx
		}(), 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Attempt(tc.ctx); got != tc.want {
				t.Errorf("Attempt() = %d, want %d", got, tc.want)
			}
		})
	}
}
