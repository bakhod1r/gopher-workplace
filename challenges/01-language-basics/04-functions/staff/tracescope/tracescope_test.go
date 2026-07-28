package tracescope

import (
	"reflect"
	"testing"
)

func TestTrace(t *testing.T) {
	got := Trace(2)
	want := []string{"start0", "end0", "start1", "end1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("=%v want %v", got, want)
	}
}
