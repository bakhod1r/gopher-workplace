package invoiceguard

import (
	"strings"
	"testing"
)

// renderer panics for any ID containing "BAD".
func renderer(id string) string {
	if strings.Contains(id, "BAD") {
		panic("nil line items for " + id)
	}
	return "doc:" + id
}

func TestRenderInvoices(t *testing.T) {
	cases := []struct {
		name    string
		ids     []string
		wantErr []bool
	}{
		{"all_render", []string{"INV-1", "INV-2"}, []bool{false, false}},
		{"one_panics", []string{"INV-1", "INV-BAD", "INV-3"}, []bool{false, true, false}},
		{"all_panic", []string{"INV-BAD", "BAD-2"}, []bool{true, true}},
		{"first_panics_rest_survive", []string{"BAD-0", "INV-1", "INV-2"}, []bool{true, false, false}},
		{"single_ok", []string{"INV-9"}, []bool{false}},
		{"empty", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RenderInvoices(tc.ids, renderer)
			if len(got) != len(tc.ids) {
				t.Fatalf("len = %d, want %d", len(got), len(tc.ids))
			}
			for i, r := range got {
				if r.ID != tc.ids[i] {
					t.Errorf("got[%d].ID = %q, want %q", i, r.ID, tc.ids[i])
				}
				if tc.wantErr[i] {
					if r.Err == nil {
						t.Errorf("got[%d].Err = nil, want a recovered panic", i)
						continue
					}
					if !strings.Contains(r.Err.Error(), tc.ids[i]) {
						t.Errorf("got[%d].Err = %q, want it to name %q", i, r.Err, tc.ids[i])
					}
					if r.Doc != "" {
						t.Errorf("got[%d].Doc = %q, want empty on failure", i, r.Doc)
					}
					continue
				}
				if r.Err != nil {
					t.Errorf("got[%d].Err = %v, want nil", i, r.Err)
				}
				if r.Doc != "doc:"+tc.ids[i] {
					t.Errorf("got[%d].Doc = %q, want %q", i, r.Doc, "doc:"+tc.ids[i])
				}
			}
		})
	}
}
