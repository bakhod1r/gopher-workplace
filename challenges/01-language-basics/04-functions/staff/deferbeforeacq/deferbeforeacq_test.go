package deferbeforeacq

import (
	"reflect"
	"testing"
)

func TestUse(t *testing.T) {
	if got := Use(false); len(got) != 0 {
		t.Errorf("nothing acquired -> no log, got %v", got)
	}
	if got := Use(true); !reflect.DeepEqual(got, []string{"open", "close"}) {
		t.Errorf("=%v want [open close]", got)
	}
}
