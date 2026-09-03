package walkchain

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	collect := func(err error) []string {
		var out []string
		Walk(err, func(e error) { out = append(out, e.Error()) })
		return out
	}

	t.Run("nil", func(t *testing.T) {
		if got := collect(nil); got != nil {
			t.Errorf("visits = %v, want none", got)
		}
	})

	t.Run("leaf", func(t *testing.T) {
		if got := collect(ErrA); !reflect.DeepEqual(got, []string{"a"}) {
			t.Errorf("visits = %v, want [a]", got)
		}
	})

	t.Run("wrapped", func(t *testing.T) {
		got := collect(fmt.Errorf("x: %w", ErrA))
		want := []string{"x: a", "a"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("visits = %v, want %v", got, want)
		}
	})

	t.Run("joined_visits_node_first", func(t *testing.T) {
		got := collect(errors.Join(ErrA, ErrB))
		if len(got) != 3 {
			t.Fatalf("visits = %v, want 3 (the join plus two leaves)", got)
		}
		if got[1] != "a" || got[2] != "b" {
			t.Errorf("branch order = %v, want a then b", got[1:])
		}
	})
}
