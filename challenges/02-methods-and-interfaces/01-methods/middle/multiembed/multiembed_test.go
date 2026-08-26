package multiembed

import "testing"

func TestCollisionName(t *testing.T) {
	c := Collision{}
	want := "B"
	if got := c.Name(); got != want {
		t.Errorf("Collision.Name() = %q, want %q", got, want)
	}
}
