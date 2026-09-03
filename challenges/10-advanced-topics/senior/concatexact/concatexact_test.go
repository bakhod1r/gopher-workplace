package concatexact

import "testing"

var sink string

func TestConcat(t *testing.T) {
	if got := Concat([]string{"a", "bc", "d"}); got != "abcd" {
		t.Errorf("Concat = %q, want \"abcd\"", got)
	}
	if got := Concat(nil); got != "" {
		t.Errorf("Concat = %q, want empty", got)
	}
	if got := Concat([]string{"", ""}); got != "" {
		t.Errorf("Concat = %q, want empty", got)
	}
	if got := Concat([]string{"solo"}); got != "solo" {
		t.Errorf("Concat = %q, want \"solo\"", got)
	}
}

func TestConcatLong(t *testing.T) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "chunk"
	}
	got := Concat(parts)
	if len(got) != 500 {
		t.Fatalf("len = %d, want 500", len(got))
	}
	for i := 0; i < len(got); i += 5 {
		if got[i:i+5] != "chunk" {
			t.Fatalf("at %d: %q, want \"chunk\"", i, got[i:i+5])
		}
	}
}

func TestConcatAllocatesOnce(t *testing.T) {
	parts := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	if n := testing.AllocsPerRun(200, func() { sink = Concat(parts) }); n > 1 {
		t.Errorf("Concat made %v allocations, want 1: size the buffer, then wrap it", n)
	}
}

func TestConcatResultIsIndependent(t *testing.T) {
	parts := []string{"ab", "cd"}
	first := Concat(parts)
	parts[0] = "XY"
	if first != "abcd" {
		t.Errorf("first = %q, want \"abcd\"", first)
	}
}
