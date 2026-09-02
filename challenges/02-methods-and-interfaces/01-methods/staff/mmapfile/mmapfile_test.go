package mmapfile

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"testing"
)

func TestReadAtFullyInside(t *testing.T) {
	m := &Mmap{Data: []byte{10, 20, 30, 40, 50}}

	p := make([]byte, 3)
	n, err := m.ReadAt(p, 1)
	if n != 3 || err != nil {
		t.Fatalf("ReadAt = (%d, %v), want (3, nil)", n, err)
	}
	if !bytes.Equal(p, []byte{20, 30, 40}) {
		t.Errorf("p = %v, want [20 30 40]", p)
	}
}

func TestReadAtShortReadAtEnd(t *testing.T) {
	m := &Mmap{Data: []byte{10, 20, 30}}

	p := make([]byte, 4)
	n, err := m.ReadAt(p, 1)
	if n != 2 || err != io.EOF {
		t.Fatalf("ReadAt = (%d, %v), want (2, io.EOF)", n, err)
	}
	if !bytes.Equal(p[:n], []byte{20, 30}) {
		t.Errorf("p[:n] = %v, want [20 30]", p[:n])
	}
}

func TestReadAtOutOfRange(t *testing.T) {
	m := &Mmap{Data: []byte{10, 20, 30}}
	p := make([]byte, 2)

	if n, err := m.ReadAt(p, 3); n != 0 || err != io.EOF {
		t.Errorf("off at end = (%d, %v), want (0, io.EOF)", n, err)
	}
	if n, err := m.ReadAt(p, 99); n != 0 || err != io.EOF {
		t.Errorf("off past end = (%d, %v), want (0, io.EOF)", n, err)
	}
	if n, err := m.ReadAt(p, -1); n != 0 || err != io.EOF {
		t.Errorf("negative off = (%d, %v), want (0, io.EOF)", n, err)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardCopiesFromRegion(t *testing.T) {
	targets := map[string]bool{"ReadAt": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "mmapfile.go", nil, 0)
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

	if !seen["copy"] || !seen["Data"] {
		t.Logf("WARN: copy out of m.Data and report the count with the io.EOF contract")
	}
}
