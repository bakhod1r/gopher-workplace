package slugify

import "testing"

func TestSlug(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Hello, World!", "hello-world"},
		{"  Go 1.26  ", "go-1-26"},
		{"a---b", "a-b"},
		{"!!!", ""},
		{"Already-Slug", "already-slug"},
	}
	for _, c := range cases {
		if got := Slug(c.in); got != c.want {
			t.Errorf("Slug(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
