package tablecaseptrbug

import "testing"

func TestPointersAreDistinct(t *testing.T) {
	cases := []Case[int]{{"a", 1}, {"b", 2}, {"c", 3}}
	got := Pointers(cases)
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for i, want := range []string{"a", "b", "c"} {
		if got[i].Name != want {
			t.Errorf("got[%d].Name = %q, want %q", i, got[i].Name, want)
		}
	}
}

func TestPointersAddressTheCases(t *testing.T) {
	cases := []Case[string]{{"a", "x"}, {"b", "y"}}
	got := Pointers(cases)
	got[0].In = "z"
	if cases[0].In != "z" {
		t.Errorf("cases[0].In = %q, want \"z\"", cases[0].In)
	}
	if cases[1].In != "y" {
		t.Errorf("cases[1].In = %q, want \"y\"", cases[1].In)
	}
}

func TestPointersEmpty(t *testing.T) {
	if got := Pointers[int](nil); len(got) != 0 {
		t.Errorf("Pointers = %v, want []", got)
	}
}
