package taglookup

import "testing"

type row struct {
	ID   int    `json:"id" db:"row_id"`
	Name string `json:"name"`
	Skip string
}

func TestTag(t *testing.T) {
	cases := []struct {
		field, key, want string
		ok               bool
	}{
		{"ID", "json", "id", true},
		{"ID", "db", "row_id", true},
		{"Name", "json", "name", true},
		{"Name", "db", "", false},
		{"Skip", "json", "", false},
		{"Missing", "json", "", false},
	}
	for _, c := range cases {
		got, ok := Tag(row{}, c.field, c.key)
		if got != c.want || ok != c.ok {
			t.Errorf("Tag(%q, %q) = %q, %v, want %q, %v", c.field, c.key, got, ok, c.want, c.ok)
		}
	}
}

func TestTagRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}} {
		if _, ok := Tag(in, "ID", "json"); ok {
			t.Errorf("Tag(%#v) reported ok, want false", in)
		}
	}
}
