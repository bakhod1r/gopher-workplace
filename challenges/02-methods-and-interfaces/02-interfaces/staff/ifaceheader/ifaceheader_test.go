package ifaceheader

import "testing"

type myError struct{}

func (*myError) Error() string { return "boom" }

func TestClassifyNil(t *testing.T) {
	if got := Classify(nil); got != "nil" {
		t.Errorf("Classify(nil) = %q, want \"nil\"", got)
	}
	typ, data := Words(nil)
	if typ != 0 || data != 0 {
		t.Errorf("Words(nil) = %d, %d; want 0, 0", typ, data)
	}
}

func TestClassifyTypedNil(t *testing.T) {
	var p *int
	if got := Classify(p); got != "typed-nil" {
		t.Errorf("Classify((*int)(nil)) = %q, want \"typed-nil\"", got)
	}
	if !IsTypedNil(p) {
		t.Error("IsTypedNil((*int)(nil)) = false")
	}

	typ, data := Words(p)
	if typ == 0 {
		t.Error("a typed nil must carry a non-zero type word")
	}
	if data != 0 {
		t.Errorf("data word = %d, want 0", data)
	}
}

func TestClassifyValue(t *testing.T) {
	if got := Classify(42); got != "value" {
		t.Errorf("Classify(42) = %q, want \"value\"", got)
	}
	n := 7
	if got := Classify(&n); got != "value" {
		t.Errorf("Classify(&n) = %q, want \"value\"", got)
	}
	if IsTypedNil(42) {
		t.Error("IsTypedNil(42) = true")
	}
}

func TestNilErrorTrap(t *testing.T) {
	var e *myError
	var err error = e

	if err == nil {
		t.Fatal("the language says a typed nil in an interface is not nil")
	}
	if !IsTypedNil(any(err)) {
		t.Error("IsTypedNil should identify the typed-nil error")
	}
	if got := Classify(any(err)); got != "typed-nil" {
		t.Errorf("Classify = %q, want \"typed-nil\"", got)
	}
}

func TestRealErrorIsValue(t *testing.T) {
	var err error = &myError{}
	if got := Classify(any(err)); got != "value" {
		t.Errorf("Classify = %q, want \"value\"", got)
	}
	if IsTypedNil(any(err)) {
		t.Error("IsTypedNil on a real error = true")
	}
}

func TestDistinctTypesHaveDistinctTypeWords(t *testing.T) {
	var a *int
	var b *string
	ta, _ := Words(a)
	tb, _ := Words(b)
	if ta == tb {
		t.Error("*int and *string should have different type words")
	}
}

func TestNilSliceIsNotATypedNilPointer(t *testing.T) {
	// A nil slice does not fit in one word, so boxing it allocates a copy of
	// the slice header: the data word points at that header and is non-zero.
	var s []int
	if got := Classify(s); got != "value" {
		t.Errorf("Classify(nil slice) = %q, want \"value\"", got)
	}

	// A nil map is a single nil pointer, so it boxes as a typed nil.
	var m map[string]int
	if got := Classify(m); got != "typed-nil" {
		t.Errorf("Classify(nil map) = %q, want \"typed-nil\"", got)
	}
}
