package factorymeth

import "testing"

func TestFactory(t *testing.T) {
	f := StoreFactory{}

	if _, ok := f.Create("mem").(MemStore); !ok {
		t.Error("expected MemStore")
	}
	if _, ok := f.Create("disk").(DiskStore); !ok {
		t.Error("expected DiskStore")
	}
	if f.Create("unknown") != nil {
		t.Error("expected nil")
	}
}
