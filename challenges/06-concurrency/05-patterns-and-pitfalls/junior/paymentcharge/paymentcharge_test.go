package paymentcharge

import (
	"errors"
	"strings"
	"testing"
)

func TestChargeAll(t *testing.T) {
	charge := func(provider string) error {
		if strings.HasPrefix(provider, "bad-") {
			return errors.New(provider + " declined")
		}
		return nil
	}

	cases := []struct {
		name      string
		providers []string
		want      []string
	}{
		{"one_failure", []string{"ok-visa", "bad-amex"}, []string{"bad-amex declined"}},
		{"all_succeed", []string{"ok-visa", "ok-mc"}, nil},
		{"all_fail_sorted", []string{"bad-z", "bad-a"}, []string{"bad-a declined", "bad-z declined"}},
		{"mixed_five", []string{"ok-a", "bad-c", "ok-b", "bad-a", "bad-b"}, []string{"bad-a declined", "bad-b declined", "bad-c declined"}},
		{"no_providers", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ChargeAll(tc.providers, charge)
			if len(got) != len(tc.want) {
				t.Fatalf("ChargeAll(%v) = %v, want %v", tc.providers, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("ChargeAll(%v) = %v, want %v", tc.providers, got, tc.want)
				}
			}
		})
	}
}
