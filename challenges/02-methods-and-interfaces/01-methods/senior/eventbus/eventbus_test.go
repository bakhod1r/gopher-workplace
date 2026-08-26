package eventbus

import "testing"

func TestEventBus(t *testing.T) {
	b := New()
	var got1, got2 string

	b.On("user.login", func(data string) { got1 = "A:" + data })
	b.On("user.login", func(data string) { got2 = "B:" + data })

	b.Emit("user.login", "alice")

	if got1 != "A:alice" || got2 != "B:alice" {
		t.Errorf("Emit failed: got1=%q got2=%q", got1, got2)
	}

	// Emit with no listeners should not panic
	b.Emit("user.logout", "bob")
}
