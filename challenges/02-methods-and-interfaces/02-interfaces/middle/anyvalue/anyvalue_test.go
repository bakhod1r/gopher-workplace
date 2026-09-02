package anyvalue

import "testing"

func TestSetGet(t *testing.T) {
	b := NewBag()
	if b.Len() != 0 {
		t.Errorf("Len = %d, want 0", b.Len())
	}

	b.Set("a", "x")
	b.Set("n", 1)

	if b.Len() != 2 {
		t.Errorf("Len = %d, want 2", b.Len())
	}
	if s, ok := b.GetString("a"); s != "x" || !ok {
		t.Errorf("GetString(a) = %q, %v", s, ok)
	}
	if s, ok := b.GetString("n"); s != "" || ok {
		t.Errorf("GetString(n) = %q, %v, want \"\", false", s, ok)
	}
	if s, ok := b.GetString("missing"); s != "" || ok {
		t.Errorf("GetString(missing) = %q, %v, want \"\", false", s, ok)
	}
}

func TestKinds(t *testing.T) {
	b := NewBag()
	b.Set("a", "x")
	b.Set("b", "y")
	b.Set("n", 1)
	b.Set("f", 1.5)
	b.Set("t", true)

	got := b.Kinds()
	want := []string{"bool", "int", "other", "string"}
	if len(got) != len(want) {
		t.Fatalf("Kinds = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Kinds = %v, want %v", got, want)
		}
	}

	if n := len(NewBag().Kinds()); n != 0 {
		t.Errorf("empty Kinds len = %d, want 0", n)
	}
}
