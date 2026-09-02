package speaker

import "testing"

func TestSpeak(t *testing.T) {
	var s Speaker

	s = Person{Name: "Go"}
	if got := s.Speak(); got != "Hi, I'm Go" {
		t.Errorf("Person = %q, want \"Hi, I'm Go\"", got)
	}

	s = Robot{ID: 1}
	if got := s.Speak(); got != "I am robot" {
		t.Errorf("Robot = %q, want \"I am robot\"", got)
	}
}

func TestIntroduce(t *testing.T) {
	if got := Introduce(Person{Name: "Ann"}); got != "Hi, I'm Ann" {
		t.Errorf("Introduce = %q", got)
	}
	if got := Introduce(Robot{ID: 9}); got != "I am robot" {
		t.Errorf("Introduce = %q", got)
	}
	if got := Introduce(Person{}); got != "Hi, I'm " {
		t.Errorf("Introduce = %q, want \"Hi, I'm \"", got)
	}
}
