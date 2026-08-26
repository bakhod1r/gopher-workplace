package mapslice

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	cases := []struct {
		name string
		list IntList
		want []string
	}{
		{"mixed", IntList{1, 2, 3}, []string{"1", "2", "3"}},
		{"negative", IntList{-5, 0}, []string{"-5", "0"}},
		{"empty", IntList{}, []string{}},
		{"nil", nil, []string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.ToString()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ToString() = %v, want %v", got, tc.want)
			}
		})
	}
}
