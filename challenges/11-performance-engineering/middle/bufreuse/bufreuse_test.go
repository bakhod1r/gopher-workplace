package bufreuse

import (
	"testing"
)

var sink []byte

func TestEncode(t *testing.T) {
	var e Encoder
	got := string(e.Encode([]string{"a", "b"}, []string{"1", "2"}))
	if got != "a=1;b=2;" {
		t.Errorf("Encode = %q, want %q", got, "a=1;b=2;")
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

func TestEncodeOverwritesThePreviousRecord(t *testing.T) {
	var e Encoder
	e.Encode([]string{"first"}, []string{"1"})
	got := string(e.Encode([]string{"second"}, []string{"2"}))
	if got != "second=2;" {
		t.Errorf("Encode = %q, want %q — the buffer must be reset, not appended to", got, "second=2;")
	}
}

func TestEncodeReusesTheBuffer(t *testing.T) {
	var e Encoder
	names := []string{"alpha", "beta", "gamma"}
	values := []string{"1", "2", "3"}
	e.Encode(names, values) // warm up
	allocs := testing.AllocsPerRun(100, func() { sink = e.Encode(names, values) })
	if allocs != 0 {
		t.Errorf("warm Encode made %v allocations, want 0", allocs)
	}
}

func TestCloneSurvivesTheNextEncode(t *testing.T) {
	var e Encoder
	e.Encode([]string{"a"}, []string{"1"})
	saved := e.Clone()
	e.Encode([]string{"b"}, []string{"2"})
	if string(saved) != "a=1;" {
		t.Errorf("clone = %q, want %q — Clone must copy, not alias", saved, "a=1;")
	}
}
