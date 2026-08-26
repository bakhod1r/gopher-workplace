package eater

import "testing"

func TestEater(t *testing.T) {
	var e Eater = Human{Name: "Alice"}
	if got := e.Eat("pizza"); got != "Alice eats pizza" {
		t.Errorf("Eat = %q", got)
	}
}
