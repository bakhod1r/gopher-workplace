package bloomfilter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestBloom(t *testing.T) {
	f := &Filter{}
	f.Add("hello")

	if !f.MightContain("hello") {
		t.Error("should contain hello")
	}
	if f.MightContain("world") {
		t.Error("should not contain world")
	}
	if !f.MightContain("ho") {
		t.Log("false positive, which is fine in Bloom")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesBothHashes(t *testing.T) {
	targets := map[string]bool{"Add": true, "MightContain": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "bloomfilter.go", nil, 0)
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

	if !seen["hash1"] || !seen["hash2"] || !seen["bits"] {
		t.Logf("WARN: set and test both hash positions in f.bits")
	}
}

// The one guarantee a Bloom filter makes: never a false negative. Every item
// ever added must still report as possibly present, whatever else collided.
func TestNoFalseNegatives(t *testing.T) {
	f := &Filter{}

	items := make([]string, 0, 5000)
	for i := 0; i < 5000; i++ {
		s := "item-" + strconv.Itoa(i)
		items = append(items, s)
		f.Add(s)
	}

	for _, s := range items {
		if !f.MightContain(s) {
			t.Fatalf("false negative for %q", s)
		}
	}
}

// The filter is a fixed [256]bool: adding items must never allocate.
func TestAddIsAllocationFree(t *testing.T) {
	f := &Filter{}

	allocs := testing.AllocsPerRun(1000, func() {
		f.Add("hello")
		f.MightContain("hello")
	})
	if allocs != 0 {
		t.Errorf("Add+MightContain allocated %.0f times per run, want 0", allocs)
	}
}
