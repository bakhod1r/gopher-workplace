package slicesminmaxfunc

import "testing"

func TestCheapestAndPriciest(t *testing.T) {
	items := []Item{{"a", 30}, {"b", 10}, {"c", 20}}
	if got, ok := Cheapest(items); !ok || got.Name != "b" {
		t.Errorf("Cheapest = %+v, %v, want b, true", got, ok)
	}
	if got, ok := Priciest(items); !ok || got.Name != "a" {
		t.Errorf("Priciest = %+v, %v, want a, true", got, ok)
	}
}

func TestEmptyDoesNotPanic(t *testing.T) {
	if got, ok := Cheapest(nil); ok || got.Name != "" {
		t.Errorf("Cheapest(nil) = %+v, %v, want zero, false", got, ok)
	}
	if got, ok := Priciest([]Item{}); ok || got.Name != "" {
		t.Errorf("Priciest(empty) = %+v, %v, want zero, false", got, ok)
	}
}

func TestSingleItem(t *testing.T) {
	items := []Item{{"only", 5}}
	lo, _ := Cheapest(items)
	hi, _ := Priciest(items)
	if lo.Name != "only" || hi.Name != "only" {
		t.Errorf("single item: %+v, %+v", lo, hi)
	}
}
