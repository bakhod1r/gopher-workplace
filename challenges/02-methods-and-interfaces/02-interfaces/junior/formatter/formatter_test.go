package formatter

import (
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	fs := []Formatter{
		Name{"John", "Doe"},
		Name{"Jane", "Smith"},
	}
	got := FormatAll(fs)
	want := []string{"Doe, John", "Smith, Jane"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FormatAll = %v, want %v", got, want)
	}
}
