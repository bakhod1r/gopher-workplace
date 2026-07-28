package handlerbug

import "testing"

func TestLabelers(t *testing.T) {
	fs := Labelers([]string{"a", "b", "c"})
	got := []string{fs[0](), fs[1](), fs[2]()}
	want := []string{"a", "b", "c"}
	for k := range want {
		if got[k] != want[k] {
			t.Errorf("closure %d = %q want %q (all captured same i?)", k, got[k], want[k])
		}
	}
}
