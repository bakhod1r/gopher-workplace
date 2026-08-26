package iterpatt

import "testing"

func TestIter(t *testing.T) {
	head := &Node{1, &Node{2, &Node{3, nil}}}
	it := NewIter(head)

	var got []int
	for it.HasNext() {
		got = append(got, it.Next())
	}

	want := []int{1, 2, 3}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("got %v, want %v", got, want)
	}
}
