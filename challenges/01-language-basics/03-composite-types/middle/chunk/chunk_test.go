package chunk

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	cases := []struct {
		xs   []int
		size int
		want [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2, 3}, 5, [][]int{{1, 2, 3}}},
		{[]int{}, 2, [][]int{}},
		{[]int{1}, 0, [][]int{}},
	}
	for _, c := range cases {
		got := Chunk(c.xs, c.size)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Chunk(%v,%d)=%v; want %v", c.xs, c.size, got, c.want)
		}
	}
}
