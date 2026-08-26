package facadepatt

import "testing"

func TestFacade(t *testing.T) {
	f := &Facade{}
	if got := f.Operation(); got != "1+2" {
		t.Errorf("Operation() = %q, want 1+2", got)
	}
}
