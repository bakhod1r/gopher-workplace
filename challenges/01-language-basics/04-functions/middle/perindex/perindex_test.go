package perindex

import "testing"

func TestHandlers(t *testing.T) {
	hs := Handlers(3)
	for i, h := range hs {
		if got := h(); got != i {
			t.Errorf("handler %d returned %d", i, got)
		}
	}
}
