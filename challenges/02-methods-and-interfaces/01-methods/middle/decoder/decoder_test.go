package decoder

import (
	"reflect"
	"testing"
)

func TestCSVDecoder(t *testing.T) {
	d := NewCSVDecoder([]string{"name,age", "Alice,30", "Bob,25"})

	expected := [][]string{
		{"name", "age"},
		{"Alice", "30"},
		{"Bob", "25"},
	}

	for i, want := range expected {
		if !d.Next() {
			t.Fatalf("row %d: Next() = false, want true", i)
		}
		if got := d.Fields(); !reflect.DeepEqual(got, want) {
			t.Errorf("row %d: Fields() = %v, want %v", i, got, want)
		}
	}

	if d.Next() {
		t.Error("extra Next() should be false")
	}

	t.Run("empty", func(t *testing.T) {
		d2 := NewCSVDecoder(nil)
		if d2.Next() {
			t.Error("empty decoder: Next() should be false")
		}
	})
}
