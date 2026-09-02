package compositepatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

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

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardRecursesIntoSub(t *testing.T) {
	targets := map[string]bool{"Size": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "compositepatt.go", nil, 0)
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

	if !seen["Sub"] || !seen["Size"] {
		t.Logf("WARN: recurse into f.Sub with Size() - summing only f.Files misses the subfolders")
	}
}

// A wide, deep tree must be summed in one traversal.
func TestLargeTree(t *testing.T) {
	const (
		depth  = 1000
		perDir = 100
	)

	leaf := func() *Folder {
		files := make([]int, perDir)
		for i := range files {
			files[i] = 1
		}
		return &Folder{Files: files}
	}

	root := leaf()
	curr := root
	for i := 1; i < depth; i++ {
		next := leaf()
		curr.Sub = append(curr.Sub, next)
		curr = next
	}

	if got, want := root.Size(), depth*perDir; got != want {
		t.Errorf("Size() = %d, want %d", got, want)
	}
}

// Summing reads the tree; it must not build an intermediate list of nodes.
func TestSizeIsAllocationFree(t *testing.T) {
	root := &Folder{
		Files: []int{10, 20},
		Sub:   []*Folder{{Files: []int{30}}, {Files: []int{40, 50}}},
	}

	allocs := testing.AllocsPerRun(1000, func() {
		root.Size()
	})
	if allocs != 0 {
		t.Errorf("Size allocated %.0f times per run, want 0", allocs)
	}
}
