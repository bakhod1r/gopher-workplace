package slicesrepeatstd

import (
	"reflect"
	"testing"
)

func TestTile(t *testing.T) {
	if got := Tile([]int{1, 2}, 2); !reflect.DeepEqual(got, []int{1, 2, 1, 2}) {
		t.Errorf("Tile = %v, want [1 2 1 2]", got)
	}
	if got := Tile([]string{"a"}, 3); !reflect.DeepEqual(got, []string{"a", "a", "a"}) {
		t.Errorf("Tile = %v, want [a a a]", got)
	}
	if got := Tile([]int{1}, 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Tile = %v, want [1]", got)
	}
}

func TestTileGuards(t *testing.T) {
	if got := Tile([]int{1}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Tile(0) = %v, want []", got)
	}
	if got := Tile([]int{1}, -1); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Tile(-1) = %v, want [] (must not panic)", got)
	}
}
