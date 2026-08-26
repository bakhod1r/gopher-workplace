package swimmer

import "testing"

func TestSwimmer(t *testing.T) {
	var s Swimmer
	s = Fish{Name: "nemo"}
	if got := s.Swim(); got != "nemo swims" { t.Errorf("Fish = %q", got) }
	s = Duck{}
	if got := s.Swim(); got != "duck swims" { t.Errorf("Duck = %q", got) }
}
