package jsonunmarshal

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name  string
		json  string
		want  int
		errOK bool
	}{
		{"dollars_and_cents", `"$10.50"`, 1050, true},
		{"cents_only", `"$0.99"`, 99, true},
		{"zero", `"$0.00"`, 0, true},
		{"exact_dollar", `"$5.00"`, 500, true},
		{"invalid_format", `"10.50"`, 0, false}, // missing $
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var m Money
			err := json.Unmarshal([]byte(tc.json), &m)
			if (err == nil) != tc.errOK {
				t.Fatalf("Unmarshal error: %v, want success: %v", err, tc.errOK)
			}
			if tc.errOK && m.Cents != tc.want {
				t.Errorf("Unmarshal %s = Money{%d}, want Money{%d}", tc.json, m.Cents, tc.want)
			}
		})
	}
}
