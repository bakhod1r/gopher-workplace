package methodval

import "testing"

func TestApplyMethod(t *testing.T) {
	cases := []struct {
		name   string
		prefix string
		arg    string
		want   string
	}{
		{"hello_alice", "Hello", "Alice", "Hello, Alice!"},
		{"hi_bob", "Hi", "Bob", "Hi, Bob!"},
		{"hey_empty", "Hey", "", "Hey, !"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fn := ApplyMethod(Greeter{Prefix: tc.prefix})
			if got := fn(tc.arg); got != tc.want {
				t.Errorf("ApplyMethod(%q)(%q) = %q, want %q",
					tc.prefix, tc.arg, got, tc.want)
			}
		})
	}

	// Ensure the function captures the receiver at creation time.
	t.Run("captures_receiver", func(t *testing.T) {
		g := Greeter{Prefix: "A"}
		fn := ApplyMethod(g)
		g.Prefix = "B" // mutating g after capture should NOT affect fn
		if got := fn("X"); got != "A, X!" {
			t.Errorf("bound method should capture at creation: got %q", got)
		}
	})
}
