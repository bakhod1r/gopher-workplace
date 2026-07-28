package gradebook

import (
	"reflect"
	"testing"
)

func TestAverages(t *testing.T) {
	book := map[string][]int{
		"ann": {90, 80, 100},
		"bob": {70, 75},
		"cid": {},
	}
	got := Averages(book)
	want := map[string]int{"ann": 90, "bob": 72}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Averages=%v; want %v", got, want)
	}
}
