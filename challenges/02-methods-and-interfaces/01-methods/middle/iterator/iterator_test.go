package iterator

import "testing"

func TestIterator(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		it := NewIntIter([]int{10, 20, 30})
		var got []int
		for it.Next() {
			got = append(got, it.Value())
		}
		want := []int{10, 20, 30}
		if len(got) != len(want) {
			t.Fatalf("len = %d, want %d", len(got), len(want))
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, got[i], want[i])
			}
		}
	})

	t.Run("empty", func(t *testing.T) {
		it := NewIntIter(nil)
		if it.Next() {
			t.Error("Next() on empty should be false")
		}
	})

	t.Run("single", func(t *testing.T) {
		it := NewIntIter([]int{42})
		if !it.Next() {
			t.Fatal("Next() should be true")
		}
		if v := it.Value(); v != 42 {
			t.Errorf("Value() = %d, want 42", v)
		}
		if it.Next() {
			t.Error("second Next() should be false")
		}
	})
}
