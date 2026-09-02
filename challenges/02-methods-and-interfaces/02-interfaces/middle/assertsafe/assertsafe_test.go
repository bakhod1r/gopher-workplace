package assertsafe

import "testing"

func TestIntHandler(t *testing.T) {
	h := &IntHandler{}
	if !h.Handle(3) || h.Sum != 3 {
		t.Errorf("Handle(3): Sum = %d", h.Sum)
	}
	if h.Handle("3") {
		t.Error("Handle(\"3\") = true, want false")
	}
	if h.Sum != 3 {
		t.Errorf("rejected payload changed Sum: %d", h.Sum)
	}
}

func TestTextHandler(t *testing.T) {
	h := &TextHandler{}
	if !h.Handle("a") {
		t.Error("Handle(\"a\") = false")
	}
	if h.Handle(1) {
		t.Error("Handle(1) = true, want false")
	}
	if len(h.Seen) != 1 || h.Seen[0] != "a" {
		t.Errorf("Seen = %v, want [a]", h.Seen)
	}
}

func TestDispatch(t *testing.T) {
	intH, textH := &IntHandler{}, &TextHandler{}
	hs := []Handler{intH, textH}

	if got := Dispatch(hs, "x"); got != 1 {
		t.Errorf("Dispatch(string) = %d, want 1", got)
	}
	if got := Dispatch(hs, 7); got != 1 {
		t.Errorf("Dispatch(int) = %d, want 1", got)
	}
	if got := Dispatch(hs, 1.5); got != 0 {
		t.Errorf("Dispatch(float) = %d, want 0", got)
	}
	if intH.Sum != 7 || len(textH.Seen) != 1 {
		t.Errorf("state: Sum=%d Seen=%v", intH.Sum, textH.Seen)
	}
}
