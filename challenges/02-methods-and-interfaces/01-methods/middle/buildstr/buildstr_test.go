package buildstr

import "testing"

func TestBuild(t *testing.T) {
	cases := []struct {
		name  string
		sep   string
		parts []string
		want  string
	}{
		{"comma", ", ", []string{"a", "b", "c"}, "a, b, c"},
		{"dash", "-", []string{"2024", "01", "01"}, "2024-01-01"},
		{"empty_parts", "|", nil, ""},
		{"single", "/", []string{"only"}, "only"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBuilder(tc.sep)
			for _, p := range tc.parts {
				b.Add(p)
			}
			if got := b.Build(); got != tc.want {
				t.Errorf("Build() = %q, want %q", got, tc.want)
			}
		})
	}

	// Test chaining.
	t.Run("chain", func(t *testing.T) {
		got := NewBuilder(" ").Add("hello").Add("world").Build()
		want := "hello world"
		if got != want {
			t.Errorf("chained Build() = %q, want %q", got, want)
		}
	})
}
