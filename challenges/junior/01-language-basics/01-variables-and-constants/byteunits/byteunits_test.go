package byteunits

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// TestNoHardcodedConstants checks the task's "no hand-typed magic numbers" rule:
// the constants should be built from iota. It does NOT fail the submission — a
// correct-but-hand-typed answer still passes. Instead it emits a `WARN:` line
// (surfaced by the runner as a warning) when no `iota` appears in the code. The
// source is parsed without comments so a mention of "iota" in a comment does not
// count.
func TestNoHardcodedConstants(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "byteunits.go", nil, 0) // 0 = skip comments
	if err != nil {
		return // parse trouble is not this check's concern
	}
	used := false
	ast.Inspect(f, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "iota" {
			used = true
		}
		return true
	})
	if !used {
		t.Logf("WARN: KiB/MiB/GiB look hand-typed — derive them from iota instead of magic numbers")
	}
}

func TestUnitConstants(t *testing.T) {
	if KiB != 1024 {
		t.Errorf("KiB = %d, want 1024", KiB)
	}
	if MiB != 1024*1024 {
		t.Errorf("MiB = %d, want %d", MiB, 1024*1024)
	}
	if GiB != 1024*1024*1024 {
		t.Errorf("GiB = %d, want %d", GiB, 1024*1024*1024)
	}
}

func TestBytes(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1024},
		{"four", 4, 4096},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Bytes(tc.n); got != tc.want {
				t.Errorf("Bytes(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
