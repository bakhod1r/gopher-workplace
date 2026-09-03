package builder

import "testing"

func TestJoin(t *testing.T) {
	if got := Join([]string{"a", "b", "c"}, "-"); got != "a-b-c" {
		t.Errorf("Join = %q, want \"a-b-c\"", got)
	}
	if got := Join([]string{"solo"}, ","); got != "solo" {
		t.Errorf("Join = %q, want \"solo\"", got)
	}
	if got := Join(nil, ","); got != "" {
		t.Errorf("Join = %q, want empty", got)
	}
	if got := Join([]string{"a", "b"}, ""); got != "ab" {
		t.Errorf("Join = %q, want \"ab\"", got)
	}
}

func TestJoinAllocationsAreBounded(t *testing.T) {
	parts := make([]string, 64)
	for i := range parts {
		parts[i] = "chunk"
	}
	if n := testing.AllocsPerRun(50, func() { _ = Join(parts, ", ") }); n > 3 {
		t.Errorf("Join made %v allocations, want at most 3: grow the buffer once", n)
	}
}
