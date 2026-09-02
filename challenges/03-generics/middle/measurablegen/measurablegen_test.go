package measurablegen

import "testing"

func TestHeaviest(t *testing.T) {
	got, ok := Heaviest([]Reading{{1}, {5}, {3}})
	if !ok || got.V != 5 {
		t.Errorf("Heaviest = %+v, %v, want {5}, true", got, ok)
	}
}

func TestHeaviestTie(t *testing.T) {
	rs := []Reading{{2}, {2}}
	got, _ := Heaviest(rs)
	if got != rs[0] {
		t.Errorf("Heaviest = %+v, want the first element", got)
	}
}

func TestHeaviestEmpty(t *testing.T) {
	got, ok := Heaviest([]Reading{})
	if ok || got.V != 0 {
		t.Errorf("Heaviest(empty) = %+v, %v, want zero, false", got, ok)
	}
}
