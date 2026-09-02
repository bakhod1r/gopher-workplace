package triesetbug

import "testing"

func TestTrieContainsInserted(t *testing.T) {
	var tr Trie[string]
	tr.Insert([]string{"a", "b"})
	if !tr.Contains([]string{"a", "b"}) {
		t.Error(`Contains([a b]) = false, want true`)
	}
}

func TestTriePrefixIsNotAMember(t *testing.T) {
	var tr Trie[string]
	tr.Insert([]string{"a", "b"})
	if tr.Contains([]string{"a"}) {
		t.Error(`Contains([a]) = true, want false`)
	}
	if tr.Contains(nil) {
		t.Error("Contains(nil) = true, want false")
	}
}

func TestTrieMissing(t *testing.T) {
	var tr Trie[string]
	tr.Insert([]string{"a"})
	if tr.Contains([]string{"z"}) {
		t.Error(`Contains([z]) = true, want false`)
	}
}
