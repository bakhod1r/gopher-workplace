package nilvsempty

import (
	"encoding/json"
	"testing"
)

func TestNonEmpty(t *testing.T) {
	out := NonEmpty([]string{"", ""})
	if out == nil {
		t.Fatal("must be non-nil empty slice")
	}
	b, _ := json.Marshal(out)
	if string(b) != "[]" {
		t.Errorf("JSON=%s; want []", b)
	}
	got := NonEmpty([]string{"a", "", "b"})
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("got %v", got)
	}
}
