package sortstructs

import (
	"reflect"
	"testing"
)

func TestByAge(t *testing.T) {
	in := []Person{{"bob", 30}, {"amy", 25}, {"cid", 30}}
	got := ByAge(in)
	want := []Person{{"amy", 25}, {"bob", 30}, {"cid", 30}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ByAge=%v; want %v", got, want)
	}
	if in[0].Name != "bob" {
		t.Error("input mutated")
	}
}
