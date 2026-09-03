package tolerantparse

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		got, err := Parse(nil)
		if got != nil || err != nil {
			t.Errorf("Parse(nil) = %v, %v, want nil, nil", got, err)
		}
	})

	t.Run("all_valid", func(t *testing.T) {
		got, err := Parse([]string{"1", "2"})
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if !reflect.DeepEqual(got, []int{1, 2}) {
			t.Errorf("got = %v, want [1 2]", got)
		}
	})

	t.Run("keeps_valid_rows", func(t *testing.T) {
		got, err := Parse([]string{"1", "x", "3"})
		if !reflect.DeepEqual(got, []int{1, 3}) {
			t.Errorf("got = %v, want [1 3]", got)
		}
		if err == nil {
			t.Fatal("err = nil, want a failure for line 1")
		}
		if !strings.Contains(err.Error(), "line 1") {
			t.Errorf("message = %q, want it to name line 1", err.Error())
		}
		if !errors.Is(err, strconv.ErrSyntax) {
			t.Error("errors.Is(err, strconv.ErrSyntax) = false, want true")
		}
	})

	t.Run("multiple_failures", func(t *testing.T) {
		_, err := Parse([]string{"x", "1", "y"})
		if err == nil {
			t.Fatal("err = nil, want failures")
		}
		if !strings.Contains(err.Error(), "line 0") || !strings.Contains(err.Error(), "line 2") {
			t.Errorf("message = %q, want it to name lines 0 and 2", err.Error())
		}
	})
}
