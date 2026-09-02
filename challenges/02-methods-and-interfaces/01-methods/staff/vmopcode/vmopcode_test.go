package vmopcode

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
	"time"
)

func TestPushAndAdd(t *testing.T) {
	v := &VM{Prog: []int{OpPush, 2, OpPush, 3, OpAdd, OpHalt}}
	if got := v.Run(); got != 5 {
		t.Errorf("Run() = %d, want 5", got)
	}
	if !reflect.DeepEqual(v.Stack, []int{5}) {
		t.Errorf("Stack = %v, want [5]", v.Stack)
	}
}

func TestDupAndMul(t *testing.T) {
	// 4, dup, mul  =>  16
	v := &VM{Prog: []int{OpPush, 4, OpDup, OpMul, OpHalt}}
	if got := v.Run(); got != 16 {
		t.Errorf("Run() = %d, want 16", got)
	}
}

func TestHaltAndEndOfProgram(t *testing.T) {
	v := &VM{Prog: []int{OpPush, 1, OpHalt, OpPush, 99}}
	if got := v.Run(); got != 1 {
		t.Errorf("Run() = %d, want 1 (instructions after HALT must not run)", got)
	}

	// Falling off the end stops just as cleanly as OpHalt.
	v2 := &VM{Prog: []int{OpPush, 7}}
	if got := v2.Run(); got != 7 {
		t.Errorf("Run() = %d, want 7", got)
	}

	// An empty program halts immediately with an empty stack.
	v3 := &VM{}
	if got := v3.Run(); got != 0 {
		t.Errorf("Run() = %d, want 0", got)
	}
}

func TestUnderflowIsANoOp(t *testing.T) {
	v := &VM{Prog: []int{OpAdd, OpPush, 3, OpHalt}}
	if got := v.Run(); got != 3 {
		t.Errorf("Run() = %d, want 3", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardDecodesProgram(t *testing.T) {
	targets := map[string]bool{"Step": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "vmopcode.go", nil, 0)
	if err != nil {
		return // parse trouble is not this check's concern
	}

	seen := map[string]bool{}
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil || !targets[fn.Name.Name] {
			continue
		}
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.SelectorExpr:
				seen[v.Sel.Name] = true
			case *ast.Ident:
				seen[v.Name] = true
			}
			return true
		})
	}

	if !seen["Prog"] || !seen["IP"] || !seen["Stack"] {
		t.Logf("WARN: fetch from v.Prog through v.IP and operate on v.Stack")
	}
}

// A long program must execute in one pass with a bounded stack: the VM may not
// leave operands behind, or a hot loop would grow without limit.
func TestLongProgramRunsInTime(t *testing.T) {
	const n = 200000

	prog := make([]int, 0, 3*n+1)
	prog = append(prog, OpPush, 0)
	for i := 0; i < n; i++ {
		prog = append(prog, OpPush, 1, OpAdd)
	}
	prog = append(prog, OpHalt)

	v := &VM{Prog: prog}

	start := time.Now()
	got := v.Run()
	elapsed := time.Since(start)

	if got != n {
		t.Fatalf("Run() = %d, want %d", got, n)
	}
	if len(v.Stack) != 1 {
		t.Errorf("stack depth = %d, want 1 — operands were left behind", len(v.Stack))
	}
	if elapsed > time.Second {
		t.Errorf("took %v, want under 1s", elapsed)
	}
}

// Stepping is pure state movement over an existing program.
func TestStepDoesNotReallocate(t *testing.T) {
	v := &VM{Prog: []int{OpPush, 1, OpPush, 2, OpAdd, OpHalt}, Stack: make([]int, 0, 8)}

	allocs := testing.AllocsPerRun(1000, func() {
		v.IP = 0
		v.Stack = v.Stack[:0]
		v.Run()
	})
	if allocs != 0 {
		t.Errorf("Run allocated %.0f times per run, want 0", allocs)
	}
}
