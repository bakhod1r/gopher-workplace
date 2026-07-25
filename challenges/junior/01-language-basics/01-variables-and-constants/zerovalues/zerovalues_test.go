package zerovalues

import "testing"

func TestDefaultConfig(t *testing.T) {
	got := DefaultConfig()

	if got.Port != 8080 {
		t.Errorf("Port = %d, want 8080", got.Port)
	}
	if got.Host != "" {
		t.Errorf("Host = %q, want \"\" (zero value)", got.Host)
	}
	if got.Debug != false {
		t.Errorf("Debug = %v, want false (zero value)", got.Debug)
	}
	if got.Tags != nil {
		t.Errorf("Tags = %v, want nil (zero value, not an empty slice)", got.Tags)
	}
	if len(got.Tags) != 0 {
		t.Errorf("len(Tags) = %d, want 0", len(got.Tags))
	}
}

func TestDefaultConfigIsFresh(t *testing.T) {
	// Mutating one returned Config must not leak into the next.
	a := DefaultConfig()
	a.Host = "example.com"
	a.Tags = append(a.Tags, "mutated")

	b := DefaultConfig()
	if b.Host != "" || b.Tags != nil {
		t.Errorf("second DefaultConfig() carried state: Host=%q Tags=%v", b.Host, b.Tags)
	}
}
