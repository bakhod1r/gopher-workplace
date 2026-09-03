package walkfields

import "testing"

type inner struct {
	M      int
	hidden int
}

type outer struct {
	N     int
	In    inner
	Ptr   *inner
	List  []inner
	Label string
}

func TestDeepSumFlat(t *testing.T) {
	if got := DeepSum(inner{M: 5}); got != 5 {
		t.Errorf("DeepSum = %d, want 5", got)
	}
}

func TestDeepSumNested(t *testing.T) {
	v := outer{
		N:    1,
		In:   inner{M: 2},
		Ptr:  &inner{M: 4},
		List: []inner{{M: 8}, {M: 16}},
	}
	if got := DeepSum(v); got != 31 {
		t.Errorf("DeepSum = %d, want 31", got)
	}
}

func TestDeepSumNilPointer(t *testing.T) {
	if got := DeepSum(outer{N: 1}); got != 1 {
		t.Errorf("DeepSum = %d, want 1: a nil pointer contributes nothing", got)
	}
}

func TestDeepSumSkipsUnexported(t *testing.T) {
	if got := DeepSum(inner{M: 1, hidden: 100}); got != 1 {
		t.Errorf("DeepSum = %d, want 1", got)
	}
}

func TestDeepSumOtherInputs(t *testing.T) {
	if got := DeepSum(7); got != 7 {
		t.Errorf("DeepSum(7) = %d, want 7", got)
	}
	if got := DeepSum([]int{1, 2, 3}); got != 6 {
		t.Errorf("DeepSum = %d, want 6", got)
	}
	if got := DeepSum(nil); got != 0 {
		t.Errorf("DeepSum(nil) = %d, want 0", got)
	}
	if got := DeepSum("x"); got != 0 {
		t.Errorf("DeepSum = %d, want 0", got)
	}
}

func TestDeepSumPointerInput(t *testing.T) {
	if got := DeepSum(&inner{M: 9}); got != 9 {
		t.Errorf("DeepSum = %d, want 9", got)
	}
}
