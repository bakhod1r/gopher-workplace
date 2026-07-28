package groupstructs

import (
	"reflect"
	"testing"
)

func TestTotalByCustomer(t *testing.T) {
	orders := []Order{
		{"ann", 100}, {"bob", 50}, {"ann", 25}, {"bob", 75},
	}
	got := TotalByCustomer(orders)
	want := map[string]int{"ann": 125, "bob": 125}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TotalByCustomer=%v; want %v", got, want)
	}
}
