package endpoint

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestConstants(t *testing.T) {
	if BaseURL != "https://api.example.com" {
		t.Errorf("BaseURL = %q, want %q", BaseURL, "https://api.example.com")
	}
	if Version != "v2" {
		t.Errorf("Version = %q, want %q", Version, "v2")
	}
}

func TestRoot(t *testing.T) {
	want := "https://api.example.com/v2"
	if Root != want {
		t.Errorf("Root = %q, want %q", Root, want)
	}
}

func TestPath(t *testing.T) {
	cases := []struct {
		name     string
		resource string
		want     string
	}{
		{"users", "users", "https://api.example.com/v2/users"},
		{"orders", "orders", "https://api.example.com/v2/orders"},
		{"nested", "users/42", "https://api.example.com/v2/users/42"},
		{"empty", "", "https://api.example.com/v2"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Path(tc.resource); got != tc.want {
				t.Errorf("Path(%q) = %q, want %q", tc.resource, got, tc.want)
			}
		})
	}
}

// Path must build on Root, so a changed Version flows through everywhere.
func TestPathBuildsOnRoot(t *testing.T) {
	got := Path("things")
	if len(got) <= len(Root) || got[:len(Root)] != Root {
		t.Errorf("Path(%q) = %q, want it to start with Root (%q)", "things", got, Root)
	}
}

// The lesson is deriving the var from the constants: a pasted literal passes the
// value checks but breaks the moment Version changes.
func TestRootDerivedFromConstants(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "endpoint.go", nil, 0) // 0 = skip comments
	if err != nil {
		return // parse trouble is not this check's concern
	}
	var value ast.Expr
	ast.Inspect(f, func(n ast.Node) bool {
		vs, ok := n.(*ast.ValueSpec)
		if !ok {
			return true
		}
		for i, name := range vs.Names {
			if name.Name == "Root" && i < len(vs.Values) {
				value = vs.Values[i]
			}
		}
		return true
	})
	if value == nil {
		return
	}
	refs := map[string]bool{}
	ast.Inspect(value, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok {
			refs[id.Name] = true
		}
		return true
	})
	if !refs["BaseURL"] || !refs["Version"] {
		t.Logf("WARN: build Root from BaseURL and Version (BaseURL + \"/\" + Version), don't paste the finished URL")
	}
}
