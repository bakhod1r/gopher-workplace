package freezebytes

import (
	"bytes"
	"testing"
	"unsafe"
)

func TestSnapshot(t *testing.T) {
	if got := Snapshot("hello"); !bytes.Equal(got, []byte("hello")) {
		t.Errorf("Snapshot = %q, want \"hello\"", got)
	}
	if got := Snapshot(""); len(got) != 0 {
		t.Errorf("Snapshot = %q, want empty", got)
	}
}

func TestSnapshotIsWritable(t *testing.T) {
	s := "abcd"
	b := Snapshot(s)
	if unsafe.SliceData(b) == unsafe.StringData(s) {
		t.Fatal("the result views the string's own bytes: writing to it is undefined")
	}
	b[0] = 'X'
	if s != "abcd" {
		t.Errorf("s = %q, want \"abcd\"", s)
	}
}

func TestSnapshotsAreIndependent(t *testing.T) {
	s := "shared"
	a := Snapshot(s)
	b := Snapshot(s)
	if unsafe.SliceData(a) == unsafe.SliceData(b) {
		t.Fatal("two snapshots share storage")
	}
	a[0] = 'X'
	if b[0] != 's' {
		t.Error("two snapshots share storage")
	}
}

func TestSnapshotHasRoomOfItsOwn(t *testing.T) {
	b := Snapshot("abc")
	b = append(b, 'd')
	if string(b) != "abcd" {
		t.Errorf("append gave %q, want \"abcd\"", string(b))
	}
}
