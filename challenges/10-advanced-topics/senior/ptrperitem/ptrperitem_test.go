package ptrperitem

import "testing"

var sink []*Node

func TestBuild(t *testing.T) {
	got := Build(3)
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for i, p := range got {
		if p == nil || p.ID != i {
			t.Fatalf("got[%d] = %v, want a node with ID %d", i, p, i)
		}
	}
	if got := Build(0); got != nil {
		t.Errorf("Build(0) = %v, want nil", got)
	}
	if got := Build(-1); got != nil {
		t.Errorf("Build(-1) = %v, want nil", got)
	}
}

func TestBuildNodesAreDistinct(t *testing.T) {
	got := Build(4)
	for i := 0; i < len(got); i++ {
		for j := i + 1; j < len(got); j++ {
			if got[i] == got[j] {
				t.Fatalf("nodes %d and %d are the same pointer", i, j)
			}
		}
	}
	got[0].ID = 99
	if got[1].ID != 1 {
		t.Error("writing one node changed another")
	}
}

func TestBuildIsWritable(t *testing.T) {
	got := Build(2)
	got[0].Next = got[1]
	if got[0].Next.ID != 1 {
		t.Error("the nodes are not linkable")
	}
}

func TestBuildAllocationsDoNotScaleWithN(t *testing.T) {
	n := testing.AllocsPerRun(50, func() { sink = Build(256) })
	if n > 4 {
		t.Errorf("Build made %v allocations for 256 nodes, want a handful: allocate one block", n)
	}
}
