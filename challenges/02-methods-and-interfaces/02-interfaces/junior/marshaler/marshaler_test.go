package marshaler

import "testing"

func TestMarshal(t *testing.T) {
	if got := (Point{X: 1, Y: 2}).Marshal(); got != "1,2" {
		t.Errorf("Point.Marshal = %q, want \"1,2\"", got)
	}
	if got := (Point{X: -3, Y: 0}).Marshal(); got != "-3,0" {
		t.Errorf("Point.Marshal = %q, want \"-3,0\"", got)
	}
	if got := Label("hi").Marshal(); got != "hi" {
		t.Errorf("Label.Marshal = %q, want \"hi\"", got)
	}
}

func TestMarshalAll(t *testing.T) {
	got := MarshalAll([]Marshaler{Point{X: 0, Y: 0}, Label("a")})
	want := []string{"0,0", "a"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if n := len(MarshalAll(nil)); n != 0 {
		t.Errorf("MarshalAll(nil) len = %d, want 0", n)
	}
}
