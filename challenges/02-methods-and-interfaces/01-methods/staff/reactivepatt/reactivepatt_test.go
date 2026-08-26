package reactivepatt

import (
	"reflect"
	"testing"
)

func TestStream(t *testing.T) {
	s := &Stream{Data: []int{1, 2, 3, 4}}

	s.Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 })

	if !reflect.DeepEqual(s.Data, []int{20, 40}) {
		t.Errorf("got %v", s.Data)
	}
}
