package swimmer

import "testing"

func TestSwim(t *testing.T) {
	var s Swimmer

	s = Fish{Name: "nemo"}
	if got := s.Swim(); got != "nemo swims" {
		t.Errorf("Fish = %q, want \"nemo swims\"", got)
	}

	s = Duck{}
	if got := s.Swim(); got != "duck swims" {
		t.Errorf("Duck = %q, want \"duck swims\"", got)
	}
}

func TestSwimAll(t *testing.T) {
	got := SwimAll([]Swimmer{Fish{Name: "a"}, Duck{}})
	want := []string{"a swims", "duck swims"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if n := len(SwimAll(nil)); n != 0 {
		t.Errorf("SwimAll(nil) len = %d, want 0", n)
	}
}
