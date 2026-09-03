package benchsetupallocbug

import "testing"

var sink []byte

func TestEncode(t *testing.T) {
	var e Encoder
	if got := string(e.Encode([]string{"a", "b"}, []string{"1", "2"})); got != "a=1;b=2;" {
		t.Errorf("Encode = %q, want %q", got, "a=1;b=2;")
	}
}

func TestEncodeReplacesThePreviousRecord(t *testing.T) {
	var e Encoder
	e.Encode([]string{"first"}, []string{"1"})
	if got := string(e.Encode([]string{"second"}, []string{"2"})); got != "second=2;" {
		t.Errorf("Encode = %q, want %q", got, "second=2;")
	}
}

func TestEncodeUsesTheShorterSlice(t *testing.T) {
	var e Encoder
	if got := string(e.Encode([]string{"a", "b"}, []string{"1"})); got != "a=1;" {
		t.Errorf("Encode = %q, want %q", got, "a=1;")
	}
	if got := string(e.Encode(nil, nil)); got != "" {
		t.Errorf("Encode = %q, want empty", got)
	}
}

func TestWarmEncodeDoesNotAllocate(t *testing.T) {
	var e Encoder
	names := []string{"alpha", "beta", "gamma"}
	values := []string{"1", "2", "3"}
	e.Encode(names, values) // warm up
	allocs := testing.AllocsPerRun(100, func() { sink = e.Encode(names, values) })
	if allocs != 0 {
		t.Errorf("warm Encode made %v allocations, want 0 — the buffer must be reused, not rebuilt", allocs)
	}
}

func TestEncodeGrowsWhenNeeded(t *testing.T) {
	var e Encoder
	names := make([]string, 100)
	values := make([]string, 100)
	for i := range names {
		names[i] = "a-fairly-long-name"
		values[i] = "a-fairly-long-value"
	}
	got := e.Encode(names, values)
	if len(got) != 100*(len("a-fairly-long-name")+len("a-fairly-long-value")+2) {
		t.Errorf("len = %d, want the full record", len(got))
	}
}
