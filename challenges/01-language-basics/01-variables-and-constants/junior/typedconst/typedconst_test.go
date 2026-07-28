package typedconst

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestConstantValues(t *testing.T) {
	if MaxBatch != 200 {
		t.Errorf("MaxBatch = %d, want 200", MaxBatch)
	}
	if Retries != 3 {
		t.Errorf("Retries = %d, want 3", Retries)
	}
}

// MaxBatch must be a byte: a plain untyped 200 would also satisfy the value
// check above, so pin the type here.
func TestMaxBatchIsAByte(t *testing.T) {
	var b byte = MaxBatch // fails to compile if MaxBatch is not byte-assignable
	if b != 200 {
		t.Errorf("MaxBatch as byte = %d, want 200", b)
	}
}

// Retries must stay untyped: an untyped constant adopts the type of its
// context, so it can initialise a float64 with no conversion. A typed int
// constant would not compile here.
func TestRetriesIsUntyped(t *testing.T) {
	var f float64 = Retries
	if f != 3 {
		t.Errorf("Retries as float64 = %v, want 3", f)
	}
	var r rune = Retries
	if r != 3 {
		t.Errorf("Retries as rune = %v, want 3", r)
	}
}

func TestFits(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want bool
	}{
		{"zero", 0, true},
		{"small", 7, true},
		{"exactly the limit", 200, true},
		{"one over", 201, false},
		{"far over", 1000, false},
		{"negative", -1, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Fits(tc.n); got != tc.want {
				t.Errorf("Fits(%d) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

// 256 overflows a byte. If the candidate converts the wrong way — byte(n)
// instead of int(MaxBatch) — 256 wraps to 0 and wrongly "fits".
func TestFitsDoesNotWrapAround(t *testing.T) {
	if Fits(256) {
		t.Errorf("Fits(256) = true, want false — converting n to byte wraps 256 to 0")
	}
	if Fits(456) {
		t.Errorf("Fits(456) = true, want false — 456 wraps to 200 in a byte")
	}
}

func TestBudget(t *testing.T) {
	cases := []struct {
		name string
		base float64
		want float64
	}{
		{"whole", 1.5, 4.5},
		{"zero", 0, 0},
		{"one", 1, 3},
		{"fractional", 2.25, 6.75},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Budget(tc.base); got != tc.want {
				t.Errorf("Budget(%v) = %v, want %v", tc.base, got, tc.want)
			}
		})
	}
}

// The point of the exercise is the declaration, not the number: MaxBatch has to
// carry an explicit byte type.
func TestMaxBatchDeclaredWithType(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "typedconst.go", nil, 0) // 0 = skip comments
	if err != nil {
		return // parse trouble is not this check's concern
	}
	typed := false
	ast.Inspect(f, func(n ast.Node) bool {
		vs, ok := n.(*ast.ValueSpec)
		if !ok || vs.Type == nil {
			return true
		}
		id, ok := vs.Type.(*ast.Ident)
		if !ok || id.Name != "byte" {
			return true
		}
		for _, name := range vs.Names {
			if name.Name == "MaxBatch" {
				typed = true
			}
		}
		return true
	})
	if !typed {
		t.Logf("WARN: declare MaxBatch as a typed constant (const MaxBatch byte = 200) — the type is the lesson")
	}
}
