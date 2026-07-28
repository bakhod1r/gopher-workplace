package groupby

import (
	"reflect"
	"testing"
)

func TestByFirst(t *testing.T) {
	got := ByFirst([]string{"apple", "banana", "avocado", "cherry", ""})
	want := map[byte][]string{
		'a': {"apple", "avocado"},
		'b': {"banana"},
		'c': {"cherry"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ByFirst=%v; want %v", got, want)
	}
}
