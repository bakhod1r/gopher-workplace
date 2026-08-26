package gcmark

import "testing"

func TestGCMark(t *testing.T) {
	o1 := &Object{}
	o2 := &Object{}
	o1.Refs = append(o1.Refs, o2)
	o1.Mark()
	if !o2.Marked {
		t.Error("expected o2 marked")
	}
}
