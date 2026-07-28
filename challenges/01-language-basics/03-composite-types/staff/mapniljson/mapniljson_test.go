package mapniljson

import (
	"encoding/json"
	"testing"
)

func TestCounts(t *testing.T) {
	m := Counts(nil)
	if m == nil {
		t.Fatal("must be non-nil empty map")
	}
	b, _ := json.Marshal(m)
	if string(b) != "{}" {
		t.Errorf("empty JSON=%s; want {}", b)
	}
	m = Counts([]string{"a", "a", "b"})
	if m["a"] != 2 || m["b"] != 1 {
		t.Errorf("counts wrong: %v", m)
	}
}
