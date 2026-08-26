package deepclone

import "testing"

func TestClone(t *testing.T) {
	orig := Config{Name: "app", Tags: []string{"v1", "prod"}}
	clone := orig.Clone()

	// Values should match.
	if clone.Name != orig.Name {
		t.Errorf("Name = %q, want %q", clone.Name, orig.Name)
	}
	if len(clone.Tags) != len(orig.Tags) {
		t.Fatalf("Tags len = %d, want %d", len(clone.Tags), len(orig.Tags))
	}
	for i := range orig.Tags {
		if clone.Tags[i] != orig.Tags[i] {
			t.Errorf("Tags[%d] = %q, want %q", i, clone.Tags[i], orig.Tags[i])
		}
	}

	// Mutation isolation.
	clone.Tags[0] = "v2"
	if orig.Tags[0] != "v1" {
		t.Errorf("mutating clone affected original: Tags[0] = %q, want v1", orig.Tags[0])
	}

	t.Run("nil_tags", func(t *testing.T) {
		c := Config{Name: "x", Tags: nil}
		cl := c.Clone()
		if cl.Tags != nil {
			t.Errorf("nil tags should clone as nil, got %v", cl.Tags)
		}
	})
}
