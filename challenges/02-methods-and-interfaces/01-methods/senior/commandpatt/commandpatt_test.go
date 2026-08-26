package commandpatt

import "testing"

func TestCommand(t *testing.T) {
	inv := &Invoker{}
	var x int

	inv.Add(func() { x += 5 })
	inv.Add(func() { x *= 2 })

	inv.ExecuteAll()
	if x != 10 {
		t.Errorf("x = %d, want 10", x)
	}

	if len(inv.commands) != 0 {
		t.Error("queue not cleared")
	}
}
