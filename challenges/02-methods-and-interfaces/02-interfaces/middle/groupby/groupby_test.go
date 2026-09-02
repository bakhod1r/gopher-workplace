package groupby

import "testing"

func TestKeyers(t *testing.T) {
	if got := (ByFirstLetter{}).Key("apple"); got != "a" {
		t.Errorf("ByFirstLetter = %q, want \"a\"", got)
	}
	if got := (ByFirstLetter{}).Key(""); got != "" {
		t.Errorf("ByFirstLetter(\"\") = %q, want empty", got)
	}
	if got := (ByLength{}).Key("abc"); got != "3" {
		t.Errorf("ByLength = %q, want \"3\"", got)
	}
}

func TestGroupByFirstLetter(t *testing.T) {
	got := Group([]string{"apple", "avocado", "beet"}, ByFirstLetter{})
	if len(got) != 2 {
		t.Fatalf("groups = %v, want 2 buckets", got)
	}
	a := got["a"]
	if len(a) != 2 || a[0] != "apple" || a[1] != "avocado" {
		t.Errorf("bucket a = %v, want [apple avocado]", a)
	}
	if len(got["b"]) != 1 {
		t.Errorf("bucket b = %v", got["b"])
	}
}

func TestGroupByLength(t *testing.T) {
	got := Group([]string{"ab", "cd", "e"}, ByLength{})
	if len(got["2"]) != 2 || len(got["1"]) != 1 {
		t.Errorf("groups = %v", got)
	}
}

func TestGroupEmpty(t *testing.T) {
	if got := Group(nil, ByLength{}); len(got) != 0 {
		t.Errorf("Group(nil) = %v, want empty", got)
	}
}

func TestSortedKeys(t *testing.T) {
	groups := Group([]string{"beet", "apple", "cider"}, ByFirstLetter{})
	got := SortedKeys(groups)
	want := []string{"a", "b", "c"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("SortedKeys = %v, want %v", got, want)
		}
	}
	if n := len(SortedKeys(nil)); n != 0 {
		t.Errorf("SortedKeys(nil) len = %d, want 0", n)
	}
}
