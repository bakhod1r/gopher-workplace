package cartesiangen

import "testing"

func TestProduct(t *testing.T) {
	got := Product([]int{1, 2}, []string{"a"})
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].First != 1 || got[0].Second != "a" {
		t.Errorf("got[0] = %+v, want {1 a}", got[0])
	}
	if got[1].First != 2 || got[1].Second != "a" {
		t.Errorf("got[1] = %+v, want {2 a}", got[1])
	}
}

func TestProductOrderIsAMajor(t *testing.T) {
	got := Product([]int{1, 2}, []string{"a", "b"})
	if len(got) != 4 {
		t.Fatalf("len = %d, want 4", len(got))
	}
	want := []string{"1a", "1b", "2a", "2b"}
	for i, w := range want {
		if string(rune('0'+got[i].First))+got[i].Second != w {
			t.Fatalf("pair %d = %+v, want %s", i, got[i], w)
		}
	}
}

func TestProductEmpty(t *testing.T) {
	got := Product([]int{}, []string{"a"})
	if got == nil || len(got) != 0 {
		t.Errorf("Product = %v, want an empty non-nil result", got)
	}
}
