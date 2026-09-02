package longesttext

import "testing"

func TestLongest(t *testing.T) {
	if got, ok := Longest([]string{"a", "bbb", "cc"}); got != "bbb" || !ok {
		t.Errorf("Longest = %q, %v, want \"bbb\", true", got, ok)
	}
	if got, ok := Longest([]Label{"xx", "yy"}); got != Label("xx") || !ok {
		t.Errorf("Longest = %q, %v, want \"xx\", true (earlier element wins ties)", got, ok)
	}
	if got, ok := Longest([]string{}); got != "" || ok {
		t.Errorf("Longest([]string{}) = %q, %v, want \"\", false", got, ok)
	}
}
