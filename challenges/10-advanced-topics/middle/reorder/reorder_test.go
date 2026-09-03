package reorder

import (
	"testing"
	"unsafe"
)

func TestSizeMatchesTheType(t *testing.T) {
	if got := Size(); got != unsafe.Sizeof(Entry{}) {
		t.Errorf("Size = %d, want %d", got, unsafe.Sizeof(Entry{}))
	}
}

func TestEntryIsPacked(t *testing.T) {
	if got := unsafe.Sizeof(Entry{}); got > 16 {
		t.Errorf("sizeof(Entry) = %d, want at most 16: reorder the fields widest first", got)
	}
}

func TestEntryStillHasEveryField(t *testing.T) {
	e := Entry{Flag: 1, Ref: 2, Kind: 3, Seq: 4}
	if e.Flag != 1 || e.Ref != 2 || e.Kind != 3 || e.Seq != 4 {
		t.Errorf("e = %+v: the fields and their types must not change", e)
	}
}

func TestFieldTypesAreUnchanged(t *testing.T) {
	e := Entry{}
	if unsafe.Sizeof(e.Ref) != 8 || unsafe.Sizeof(e.Seq) != 4 ||
		unsafe.Sizeof(e.Flag) != 1 || unsafe.Sizeof(e.Kind) != 1 {
		t.Error("a field's type changed; only the declaration order may move")
	}
}
