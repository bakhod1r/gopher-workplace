package methodtypeparamgen

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAddChains(t *testing.T) {
	got := Bag[int]{}.Add(1).Add(2).Items()
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Items() = %v, want [1 2]", got)
	}
}

func TestAddDoesNotMutateTheReceiver(t *testing.T) {
	b := Bag[int]{}.Add(1)
	b.Add(2)
	if got := b.Items(); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Items() = %v, want [1]", got)
	}
}

func TestMapBagChangesType(t *testing.T) {
	b := Bag[int]{}.Add(1).Add(2)
	got := MapBag(b, strconv.Itoa).Items()
	if !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("MapBag = %v, want [1 2]", got)
	}
	empty := MapBag(Bag[int]{}, strconv.Itoa).Items()
	if len(empty) != 0 {
		t.Errorf("MapBag(empty) = %v, want []", empty)
	}
}
