package byteappend

import (
	"testing"
)

var sink []byte

func TestAppendRecord(t *testing.T) {
	if got := string(AppendRecord(nil, "a", "1")); got != "a=1;" {
		t.Errorf("AppendRecord = %q, want %q", got, "a=1;")
	}
	if got := string(AppendRecord([]byte("x:"), "k", "v")); got != "x:k=v;" {
		t.Errorf("AppendRecord = %q, want %q", got, "x:k=v;")
	}
	if got := string(AppendRecord(nil, "", "")); got != "=;" {
		t.Errorf("AppendRecord = %q, want %q", got, "=;")
	}
}

func TestAppendRecordChains(t *testing.T) {
	buf := AppendRecord(nil, "a", "1")
	buf = AppendRecord(buf, "b", "2")
	if got := string(buf); got != "a=1;b=2;" {
		t.Errorf("chained = %q, want %q", got, "a=1;b=2;")
	}
}

func TestAppendRecordReusesCapacity(t *testing.T) {
	buf := make([]byte, 0, 64)
	allocs := testing.AllocsPerRun(50, func() { sink = AppendRecord(buf[:0], "key", "value") })
	if allocs != 0 {
		t.Errorf("AppendRecord into a buffer with spare capacity made %v allocations, want 0", allocs)
	}
}
