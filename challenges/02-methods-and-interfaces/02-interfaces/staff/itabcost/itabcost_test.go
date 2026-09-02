package itabcost

import "testing"

func TestApply(t *testing.T) {
	if got := (AddOp{N: 2}).Apply(0, 1); got != 3 {
		t.Errorf("AddOp.Apply = %d, want 3", got)
	}
	if got := (MulOp{}).Apply(3, 4); got != 12 {
		t.Errorf("MulOp.Apply = %d, want 12", got)
	}
}

func TestRunIfaceAdd(t *testing.T) {
	if got := RunIface(AddOp{N: 2}, 0, []int{1, 2, 3}); got != 12 {
		t.Errorf("RunIface = %d, want 12", got)
	}
}

func TestRunIfaceMul(t *testing.T) {
	if got := RunIface(MulOp{}, 1, []int{2, 3, 4}); got != 24 {
		t.Errorf("RunIface = %d, want 24", got)
	}
}

func TestBothPathsAgree(t *testing.T) {
	op := AddOp{N: 3}
	vs := []int{1, 2, 3, 4, 5}
	a := RunIface(op, 10, vs)
	b := RunConcrete(op, 10, vs)
	if a != b {
		t.Errorf("RunIface = %d, RunConcrete = %d; they must agree", a, b)
	}
	if a != 40 {
		t.Errorf("result = %d, want 40", a)
	}
}

func TestEmptyInput(t *testing.T) {
	if got := RunIface(AddOp{N: 1}, 7, nil); got != 7 {
		t.Errorf("RunIface = %d, want the start value 7", got)
	}
	if got := RunConcrete(AddOp{N: 1}, 7, nil); got != 7 {
		t.Errorf("RunConcrete = %d, want 7", got)
	}
}

func TestNeitherPathAllocates(t *testing.T) {
	vs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var op Op = AddOp{N: 1}

	if avg := testing.AllocsPerRun(500, func() { _ = RunIface(op, 0, vs) }); avg > 0 {
		t.Errorf("RunIface allocated %.2f times per call, want 0", avg)
	}
	if avg := testing.AllocsPerRun(500, func() { _ = RunConcrete(AddOp{N: 1}, 0, vs) }); avg > 0 {
		t.Errorf("RunConcrete allocated %.2f times per call, want 0", avg)
	}
}

func BenchmarkRunIface(b *testing.B) {
	vs := make([]int, 1000)
	for i := range vs {
		vs[i] = i
	}
	var op Op = AddOp{N: 1}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = RunIface(op, 0, vs)
	}
}

func BenchmarkRunConcrete(b *testing.B) {
	vs := make([]int, 1000)
	for i := range vs {
		vs[i] = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = RunConcrete(AddOp{N: 1}, 0, vs)
	}
}
