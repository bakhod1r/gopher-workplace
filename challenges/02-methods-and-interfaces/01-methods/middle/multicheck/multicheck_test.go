package multicheck

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name string
		u    User
		want []string
	}{
		{"all_valid", User{"Go", "go@go.dev", 10}, nil},
		{"all_invalid", User{"", "bad", -1},
			[]string{"name is required", "invalid email", "age must be non-negative"}},
		{"empty_name", User{"", "ok@ok.com", 5}, []string{"name is required"}},
		{"bad_email", User{"X", "nope", 0}, []string{"invalid email"}},
		{"neg_age", User{"X", "x@x", -1}, []string{"age must be non-negative"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.u.Validate()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Validate() = %v, want %v", got, tc.want)
			}
		})
	}
}
