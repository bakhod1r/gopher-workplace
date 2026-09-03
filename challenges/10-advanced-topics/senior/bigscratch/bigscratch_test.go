package bigscratch

import (
	"bytes"
	"runtime"
	"testing"
)

var sink []byte

func TestFormat(t *testing.T) {
	cases := map[int64]string{
		0: "0", 42: "42", -7: "-7",
		9223372036854775807:  "9223372036854775807",
		-9223372036854775808: "-9223372036854775808",
	}
	for in, want := range cases {
		if got := Format(in); !bytes.Equal(got, []byte(want)) {
			t.Errorf("Format(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatResultsAreIndependent(t *testing.T) {
	a := Format(11)
	b := Format(22)
	if !bytes.Equal(a, []byte("11")) {
		t.Errorf("a = %q, want \"11\"", a)
	}
	if !bytes.Equal(b, []byte("22")) {
		t.Errorf("b = %q, want \"22\"", b)
	}
}

func TestFormatDoesNotAllocateAKilobyte(t *testing.T) {
	const runs = 2000
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	for i := 0; i < runs; i++ {
		sink = Format(int64(i))
	}
	runtime.ReadMemStats(&after)
	used := after.TotalAlloc - before.TotalAlloc
	if used > runs*128 {
		t.Errorf("allocated %d bytes over %d calls (%d per call), want well under 128 per call: the scratch buffer escapes",
			used, runs, used/runs)
	}
}
