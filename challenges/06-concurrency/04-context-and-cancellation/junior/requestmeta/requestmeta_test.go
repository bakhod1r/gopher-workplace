package requestmeta

import (
	"context"
	"testing"
)

func TestMeta(t *testing.T) {
	bg := context.Background()

	cases := []struct {
		name       string
		ctx        context.Context
		wantTenant string
		wantTrace  string
	}{
		{"nothing_attached", bg, "", ""},
		{"both_present", WithMeta(bg, "acme", "4bf9"), "acme", "4bf9"},
		{"empty_tenant", WithMeta(bg, "", "4bf9"), "", "4bf9"},
		{"inner_wins", WithMeta(WithMeta(bg, "a", "1"), "b", "2"), "b", "2"},
		{"only_tenant_key_set", context.WithValue(bg, tenantKey{}, "solo"), "solo", ""},
		{"survives_derivation", func() context.Context {
			ctx, cancel := context.WithCancel(WithMeta(bg, "acme", "4bf9"))
			defer cancel()
			return ctx
		}(), "acme", "4bf9"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tenant, trace := Meta(tc.ctx)
			if tenant != tc.wantTenant || trace != tc.wantTrace {
				t.Errorf("Meta() = (%q, %q), want (%q, %q)", tenant, trace, tc.wantTenant, tc.wantTrace)
			}
		})
	}
}
