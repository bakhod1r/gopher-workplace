package speaker

import "testing"

func TestSpeaker(t *testing.T) {
	var s Speaker
	s = Person{Name: "Go"}
	if got := s.Speak(); got != "Hi, I'm Go" { t.Errorf("Person = %q", got) }
	s = Robot{ID: 1}
	if got := s.Speak(); got != "I am robot" { t.Errorf("Robot = %q", got) }
}
