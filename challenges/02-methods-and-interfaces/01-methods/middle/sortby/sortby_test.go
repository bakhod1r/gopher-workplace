package sortby

import (
	"reflect"
	"testing"
)

func TestSortByField(t *testing.T) {
	cases := []struct {
		name   string
		people []Person
		want   []Person
	}{
		{
			"by_age",
			[]Person{{"Charlie", 30}, {"Alice", 20}, {"Bob", 25}},
			[]Person{{"Alice", 20}, {"Bob", 25}, {"Charlie", 30}},
		},
		{
			"already_sorted",
			[]Person{{"A", 1}, {"B", 2}},
			[]Person{{"A", 1}, {"B", 2}},
		},
		{
			"single",
			[]Person{{"Only", 5}},
			[]Person{{"Only", 5}},
		},
		{
			"empty",
			[]Person{},
			[]Person{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			SortByField(tc.people, Person.ByAge)
			if !reflect.DeepEqual(tc.people, tc.want) {
				t.Errorf("SortByField result = %v, want %v", tc.people, tc.want)
			}
		})
	}
}
