package compositepatt

import "testing"

func TestComposite(t *testing.T) {
	root := &Folder{
		Files: []int{10, 20},
		Sub: []*Folder{
			{Files: []int{30}},
			{Files: []int{40, 50}},
		},
	}
	if got := root.Size(); got != 150 {
		t.Errorf("Size() = %d, want 150", got)
	}
}
