package stringerifc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDescribe(t *testing.T) {
	got := Describe([]Tag{"a", "b"})
	if !reflect.DeepEqual(got, []string{"tag:a", "tag:b"}) {
		t.Errorf("Describe = %v, want [tag:a tag:b]", got)
	}
	if got := Describe([]Tag{}); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("Describe([]) = %v, want []", got)
	}
}

func TestDescribeMatchesDescribeAny(t *testing.T) {
	typed := Describe([]Tag{"a"})
	boxed := DescribeAny([]fmt.Stringer{Tag("a")})
	if !reflect.DeepEqual(typed, boxed) {
		t.Errorf("Describe = %v, DescribeAny = %v, want the same result", typed, boxed)
	}
}
