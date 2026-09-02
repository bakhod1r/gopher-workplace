package methodset

import "testing"

func TestNames(t *testing.T) {
	got := Names([]Named{Value{N: "a"}, &Pointer{N: "b"}})
	want := []string{"a", "b"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestRename(t *testing.T) {
	p := &Pointer{N: "old"}
	p.Rename("new")
	if p.Name() != "new" {
		t.Errorf("Name = %q, want \"new\"", p.Name())
	}
}

func TestSatisfies(t *testing.T) {
	cases := []struct {
		name string
		v    any
		want bool
	}{
		{"pointer_ptr", &Pointer{}, true},
		{"pointer_value", Pointer{}, false},
		{"value_value", Value{}, false},
		{"value_ptr", &Value{}, false},
		{"nil", nil, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Satisfies(tc.v); got != tc.want {
				t.Errorf("Satisfies(%T) = %v, want %v", tc.v, got, tc.want)
			}
		})
	}
}

func TestValueInNamedButNotPointerMethods(t *testing.T) {
	var n Named = Value{N: "x"}
	if n.Name() != "x" {
		t.Errorf("Name = %q", n.Name())
	}
	if _, ok := n.(Renamer); ok {
		t.Error("Value must not satisfy Renamer")
	}
}
