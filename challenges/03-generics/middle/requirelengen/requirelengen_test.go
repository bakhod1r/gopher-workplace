package requirelengen

import "testing"

type fataller struct {
	testing.TB
	fatals int
	helper bool
}

func (f *fataller) Helper()                           { f.helper = true }
func (f *fataller) Fatalf(format string, args ...any) { f.fatals++ }
func (f *fataller) Errorf(format string, args ...any) { f.fatals++ }

func TestRequireLenPasses(t *testing.T) {
	f := &fataller{}
	RequireLen(f, []int{1, 2}, 2)
	if f.fatals != 0 {
		t.Errorf("fatals = %d, want 0", f.fatals)
	}
	if !f.helper {
		t.Error("RequireLen did not call t.Helper()")
	}
}

func TestRequireLenFails(t *testing.T) {
	f := &fataller{}
	RequireLen(f, []int{}, 1)
	if f.fatals != 1 {
		t.Errorf("fatals = %d, want 1", f.fatals)
	}
}

func TestRequireLenStructs(t *testing.T) {
	type row struct{ tags []string }
	f := &fataller{}
	RequireLen(f, []row{{}}, 1)
	if f.fatals != 0 {
		t.Errorf("fatals = %d, want 0 (elements need no constraint)", f.fatals)
	}
}
