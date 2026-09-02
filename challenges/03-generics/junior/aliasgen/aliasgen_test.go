package aliasgen

import "testing"

func TestIndex(t *testing.T) {
	ix := NewIndex[string]()
	if ix == nil {
		t.Fatal("NewIndex returned nil, want an allocated index")
	}
	if Marked(ix, "a") {
		t.Error(`Marked(ix, "a") on a fresh index = true, want false`)
	}
	Mark(ix, "a")
	if !Marked(ix, "a") {
		t.Error(`Marked(ix, "a") = false, want true`)
	}
	if Marked(ix, "b") {
		t.Error(`Marked(ix, "b") = true, want false`)
	}
}

func TestIndexIsAnAlias(t *testing.T) {
	// Because Index[K] is an alias, a plain map is the very same type.
	plain := map[string]struct{}{}
	Mark(plain, "a")
	if !Marked(plain, "a") {
		t.Error("a plain map[string]struct{} did not work as an Index[string]")
	}
	var ix Index[int] = map[int]struct{}{1: {}}
	if !Marked(ix, 1) {
		t.Error("assigning a map literal to Index[int] failed")
	}
}
