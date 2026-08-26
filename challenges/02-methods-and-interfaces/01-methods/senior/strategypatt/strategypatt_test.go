package strategypatt

import (
	"reflect"
	"testing"
)

func TestStrategy(t *testing.T) {
	c := &Context{Data: []int{1, 2, 3}}

	double := func(x int) int { return x * 2 }
	c.Process(double)

	if !reflect.DeepEqual(c.Data, []int{2, 4, 6}) {
		t.Errorf("Process double = %v", c.Data)
	}
}
