package flyweightpatt

import "testing"

func TestFlyweight(t *testing.T) {
	f := &FlyweightFactory{fonts: make(map[string]*FontData)}

	f1 := f.Get("Arial")
	f2 := f.Get("Arial")
	f3 := f.Get("Times")

	if f1 != f2 {
		t.Error("expected same Arial instance")
	}
	if f1 == f3 {
		t.Error("expected different Times instance")
	}
}
