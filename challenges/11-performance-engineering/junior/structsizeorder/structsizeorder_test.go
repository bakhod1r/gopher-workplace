package structsizeorder

import (
	"testing"
	"unsafe"
)

func TestNewRecord(t *testing.T) {
	r := NewRecord(1, 2, 3, true)
	if r.ID != 1 || r.Count != 2 || r.Kind != 3 || !r.Enabled {
		t.Errorf("NewRecord = %+v, want ID 1, Count 2, Kind 3, Enabled true", r)
	}
	z := NewRecord(0, 0, 0, false)
	if z != (Record{}) {
		t.Errorf("NewRecord(zeros) = %+v, want the zero Record", z)
	}
}

func TestRecordFieldTypes(t *testing.T) {
	var r Record
	var (
		_ int64 = r.ID
		_ int32 = r.Count
		_ int16 = r.Kind
		_ bool  = r.Enabled
	)
}

func TestRecordIsPacked(t *testing.T) {
	if got := unsafe.Sizeof(Record{}); got != 16 {
		t.Errorf("unsafe.Sizeof(Record{}) = %d, want 16 — reorder the fields widest first", got)
	}
}

func TestRecordSliceIsPacked(t *testing.T) {
	rs := make([]Record, 1000)
	if got := unsafe.Sizeof(rs[0]) * 1000; got != 16000 {
		t.Errorf("1000 records occupy %d bytes, want 16000", got)
	}
}
