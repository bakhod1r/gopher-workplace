package visitorpatt

import "testing"

func TestVisitor(t *testing.T) {
	root := &Node{
		Val:   1,
		Left:  &Node{Val: 2},
		Right: &Node{Val: 3},
	}

	sum := 0
	root.Accept(func(v int) { sum += v })

	if sum != 6 {
		t.Errorf("sum = %d, want 6", sum)
	}
}
