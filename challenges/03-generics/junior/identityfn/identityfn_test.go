package identityfn

import "testing"

func TestIdentity(t *testing.T) {
	if got := Identity(7); got != 7 {
		t.Errorf("Identity(7) = %v, want 7", got)
	}
	if got := Identity("go"); got != "go" {
		t.Errorf("Identity(%q) = %q, want %q", "go", got, "go")
	}
	if got := Identity(true); got != true {
		t.Errorf("Identity(true) = %v, want true", got)
	}
}
