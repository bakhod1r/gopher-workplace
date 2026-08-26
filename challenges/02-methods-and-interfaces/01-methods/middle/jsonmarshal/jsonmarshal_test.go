package jsonmarshal

import (
	"encoding/json"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	cases := []struct {
		name  string
		cents int
		want  string
	}{
		{"dollars_and_cents", 1050, `"$10.50"`},
		{"cents_only", 99, `"$0.99"`},
		{"zero", 0, `"$0.00"`},
		{"exact_dollar", 500, `"$5.00"`},
		{"large", 123456, `"$1234.56"`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := Money{Cents: tc.cents}
			got, err := json.Marshal(m)
			if err != nil {
				t.Fatalf("Marshal error: %v", err)
			}
			if string(got) != tc.want {
				t.Errorf("Marshal(Money{%d}) = %s, want %s", tc.cents, got, tc.want)
			}
		})
	}
}
