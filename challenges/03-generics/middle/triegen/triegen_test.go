package triegen

import "testing"

func TestTrieInsertGet(t *testing.T) {
	tr := NewTrie[int]()
	tr.Insert("go", 1)
	tr.Insert("gopher", 2)
	if v, ok := tr.Get("go"); v != 1 || !ok {
		t.Errorf(`Get("go") = %v, %v, want 1, true`, v, ok)
	}
	if v, ok := tr.Get("gopher"); v != 2 || !ok {
		t.Errorf(`Get("gopher") = %v, %v, want 2, true`, v, ok)
	}
	if v, ok := tr.Get("g"); v != 0 || ok {
		t.Errorf(`Get("g") = %v, %v, want 0, false (prefix is not a key)`, v, ok)
	}
	if _, ok := tr.Get("java"); ok {
		t.Error(`Get("java") reported ok`)
	}
}

func TestTrieStoredZero(t *testing.T) {
	tr := NewTrie[int]()
	tr.Insert("x", 0)
	if v, ok := tr.Get("x"); v != 0 || !ok {
		t.Errorf(`Get("x") = %v, %v, want 0, true (a stored zero is still a key)`, v, ok)
	}
}

func TestTrieHasPrefix(t *testing.T) {
	tr := NewTrie[int]()
	tr.Insert("go", 1)
	if !tr.HasPrefix("g") {
		t.Error(`HasPrefix("g") = false, want true`)
	}
	if !tr.HasPrefix("") {
		t.Error(`HasPrefix("") = false, want true`)
	}
	if tr.HasPrefix("ja") {
		t.Error(`HasPrefix("ja") = true, want false`)
	}
}

func TestTrieUnicode(t *testing.T) {
	tr := NewTrie[string]()
	tr.Insert("привет", "hi")
	if v, ok := tr.Get("привет"); v != "hi" || !ok {
		t.Errorf("Get(unicode) = %q, %v, want hi, true", v, ok)
	}
	if !tr.HasPrefix("при") {
		t.Error("HasPrefix(unicode prefix) = false, want true")
	}
}
