package escapes

import "testing"

func TestRow(t *testing.T) {
	got := Row("id", "42")
	want := "id\t42\n"
	if got != want {
		t.Errorf("Row=%q; want %q", got, want)
	}
}
