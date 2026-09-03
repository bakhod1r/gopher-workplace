package resetbuf

import "testing"

func TestResetEmptiesAndKeepsCapacity(t *testing.T) {
	buf := make([]byte, 8, 64)
	out := Reset(buf)
	if len(out) != 0 {
		t.Errorf("len = %d, want 0", len(out))
	}
	if cap(out) != 64 {
		t.Errorf("cap = %d, want 64: the capacity must survive the reset", cap(out))
	}
}

func TestResetReusesTheSameArray(t *testing.T) {
	buf := make([]byte, 0, 8)
	buf = append(buf, 'a', 'b')
	out := append(Reset(buf), 'z')
	if len(out) != 1 || out[0] != 'z' {
		t.Fatalf("out = %q, want \"z\"", out)
	}
	if &buf[:1][0] != &out[0] {
		t.Error("the reset buffer allocated a new array instead of reusing buf")
	}
}

func TestResetNil(t *testing.T) {
	if got := Reset(nil); len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}
