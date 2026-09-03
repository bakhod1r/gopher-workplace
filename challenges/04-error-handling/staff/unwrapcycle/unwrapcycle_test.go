package unwrapcycle

import (
	"fmt"
	"reflect"
	"testing"
)

// selfError unwraps to itself.
type selfError struct{ msg string }

func (e *selfError) Error() string { return e.msg }
func (e *selfError) Unwrap() error { return e }

func TestChain(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := Chain(nil); got != nil {
			t.Errorf("Chain(nil) = %v, want nil", got)
		}
	})

	t.Run("normal_chain", func(t *testing.T) {
		got := Chain(fmt.Errorf("x: %w", ErrA))
		want := []string{"x: a", "a"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Chain = %v, want %v", got, want)
		}
	})

	t.Run("self_cycle_terminates", func(t *testing.T) {
		done := make(chan []string, 1)
		go func() { done <- Chain(&selfError{msg: "loop"}) }()
		select {
		case got := <-done:
			if len(got) != 1 || got[0] != "loop" {
				t.Errorf("Chain = %v, want [loop]", got)
			}
		default:
			// fall through to a blocking receive below
		}
		got := <-done
		if len(got) != 1 {
			t.Errorf("Chain = %v, want a single entry", got)
		}
	})
}
