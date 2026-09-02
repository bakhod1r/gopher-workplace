package treegen

import (
	"reflect"
	"testing"
)

func TestInsertAndInOrder(t *testing.T) {
	var root *TreeNode[int]
	for _, v := range []int{2, 1, 3} {
		root = Insert(root, v)
	}
	if got := InOrder(root); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InOrder = %v, want [1 2 3]", got)
	}
}

func TestInsertIgnoresDuplicates(t *testing.T) {
	var root *TreeNode[int]
	for _, v := range []int{2, 2, 1} {
		root = Insert(root, v)
	}
	if got := InOrder(root); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("InOrder = %v, want [1 2]", got)
	}
}

func TestInOrderNil(t *testing.T) {
	if got := InOrder[int](nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("InOrder(nil) = %v, want []", got)
	}
}

func TestInsertStrings(t *testing.T) {
	var root *TreeNode[string]
	for _, v := range []string{"b", "a", "c"} {
		root = Insert(root, v)
	}
	if got := InOrder(root); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("InOrder = %v, want [a b c]", got)
	}
}
