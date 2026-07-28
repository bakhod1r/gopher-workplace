package structfilter

import (
	"reflect"
	"testing"
)

func TestActiveNames(t *testing.T) {
	users := []User{
		{"ann", true}, {"bob", false}, {"cid", true},
	}
	got := ActiveNames(users)
	want := []string{"ann", "cid"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ActiveNames=%v; want %v", got, want)
	}
	if len(ActiveNames(nil)) != 0 {
		t.Error("nil -> empty")
	}
}
