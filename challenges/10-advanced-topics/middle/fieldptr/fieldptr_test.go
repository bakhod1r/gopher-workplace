package fieldptr

import "testing"

func TestBumpSeq(t *testing.T) {
	r := &Rec{Tag: 7, Seq: 1, Name: "n"}
	if got := BumpSeq(r); got != 2 {
		t.Errorf("BumpSeq = %d, want 2", got)
	}
	if r.Seq != 2 {
		t.Errorf("r.Seq = %d, want 2: the write must reach the caller's record", r.Seq)
	}
}

func TestBumpSeqLeavesTheOtherFields(t *testing.T) {
	r := &Rec{Tag: 7, Seq: 0, Name: "name"}
	BumpSeq(r)
	if r.Tag != 7 || r.Name != "name" {
		t.Errorf("r = %+v: only Seq may change", *r)
	}
}

func TestBumpSeqRepeated(t *testing.T) {
	r := &Rec{}
	for i := 1; i <= 100; i++ {
		if got := BumpSeq(r); got != int64(i) {
			t.Fatalf("call %d returned %d", i, got)
		}
	}
}

func TestBumpSeqNegative(t *testing.T) {
	r := &Rec{Seq: -1}
	if got := BumpSeq(r); got != 0 {
		t.Errorf("BumpSeq = %d, want 0", got)
	}
}
