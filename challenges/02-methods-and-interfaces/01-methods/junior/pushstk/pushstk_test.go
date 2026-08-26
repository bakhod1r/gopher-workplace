package pushstk

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	cases := []struct {
		name   string
		start  []int
		pushes []int
		want   []int
	}{
		{"empty_push_one", nil, []int{5}, []int{5}},
		{"push_two", nil, []int{5, 3}, []int{5, 3}},
		{"existing", []int{1, 2}, []int{3}, []int{1, 2, 3}},
		{"push_zero", nil, []int{0}, []int{0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stack{Items: tc.start}
			for _, v := range tc.pushes {
				s.Push(v)
			}
			if !reflect.DeepEqual(s.Items, tc.want) {
				t.Errorf("after Push(%v) on %v: Items = %v, want %v",
					tc.pushes, tc.start, s.Items, tc.want)
			}
		})
	}
}
