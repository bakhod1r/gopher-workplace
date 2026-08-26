package encoder

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	e := &CSVEncoder{}
	e.Encode("name", "age", "city")
	e.Encode("Alice", "30", "NYC")
	e.Encode("Bob", "25", "LA")

	want := []string{"name,age,city", "Alice,30,NYC", "Bob,25,LA"}
	if !reflect.DeepEqual(e.Rows, want) {
		t.Errorf("Rows = %v, want %v", e.Rows, want)
	}

	t.Run("single_field", func(t *testing.T) {
		e2 := &CSVEncoder{}
		e2.Encode("only")
		if !reflect.DeepEqual(e2.Rows, []string{"only"}) {
			t.Errorf("single field: Rows = %v", e2.Rows)
		}
	})

	t.Run("empty_fields", func(t *testing.T) {
		e3 := &CSVEncoder{}
		e3.Encode()
		if !reflect.DeepEqual(e3.Rows, []string{""}) {
			t.Errorf("empty fields: Rows = %v", e3.Rows)
		}
	})
}
