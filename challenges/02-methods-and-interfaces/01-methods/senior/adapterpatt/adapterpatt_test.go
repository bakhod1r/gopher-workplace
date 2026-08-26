package adapterpatt

import "testing"

func TestAdapter(t *testing.T) {
	a := &ModernAdapter{}
	if got := a.GetIntData(); got != 123 {
		t.Errorf("GetIntData() = %d, want 123", got)
	}
}
