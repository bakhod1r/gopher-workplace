package maskpanic

import "testing"

func TestFirstPanic(t *testing.T) {
	got := FirstPanic(func() { panic("original") })
	if got != "original" {
		t.Errorf("got %v want \"original\" (cleanup masked it?)", got)
	}
}
