package deepequalcost

import "testing"

var sink bool

func TestChanged(t *testing.T) {
	base := Config{Retries: 1, Timeout: 2, Name: "n", Debug: true}
	if Changed(base, base) {
		t.Error("Changed = true for identical configs, want false")
	}
	other := base
	other.Retries = 9
	if !Changed(base, other) {
		t.Error("Changed = false for differing configs, want true")
	}
}

func TestChangedEveryField(t *testing.T) {
	base := Config{}
	cases := []Config{
		{Retries: 1},
		{Timeout: 1},
		{Name: "x"},
		{Debug: true},
	}
	for i, c := range cases {
		if !Changed(base, c) {
			t.Errorf("case %d: Changed = false, want true", i)
		}
	}
}

func TestChangedZeroValues(t *testing.T) {
	if Changed(Config{}, Config{}) {
		t.Error("Changed = true for two zero configs, want false")
	}
}

func TestChangedAllocatesNothing(t *testing.T) {
	a := Config{Retries: 1, Name: "left"}
	b := Config{Retries: 2, Name: "right"}
	if n := testing.AllocsPerRun(200, func() { sink = Changed(a, b) }); n != 0 {
		t.Errorf("Changed made %v allocations, want 0: the struct is comparable", n)
	}
}
