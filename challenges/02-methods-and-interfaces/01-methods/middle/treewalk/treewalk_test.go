package treewalk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name string
		tree *Tree
		want []int
	}{
		{"nil", nil, []int{}},
		{"single", &Tree{Val: 5}, []int{5}},
		{"balanced", &Tree{2, &Tree{Val: 1}, &Tree{Val: 3}}, []int{1, 2, 3}},
		{"left_only", &Tree{3, &Tree{2, &Tree{Val: 1}, nil}, nil}, []int{1, 2, 3}},
		{"right_only", &Tree{1, nil, &Tree{2, nil, &Tree{Val: 3}}}, []int{1, 2, 3}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.tree.Walk()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Walk() = %v, want %v", got, tc.want)
			}
		})
	}
}
