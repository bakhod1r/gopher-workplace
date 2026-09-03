package fieldsview

import (
	"reflect"
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	if got := Fields("a,b,c", ','); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Fields = %q, want [a b c]", got)
	}
	if got := Fields("", ','); !reflect.DeepEqual(got, []string{""}) {
		t.Errorf("Fields = %q, want [\"\"]", got)
	}
	if got := Fields("a,,b", ','); !reflect.DeepEqual(got, []string{"a", "", "b"}) {
		t.Errorf("Fields = %q, want [a  b]", got)
	}
	if got := Fields(",x", ','); !reflect.DeepEqual(got, []string{"", "x"}) {
		t.Errorf("Fields = %q, want [ x]", got)
	}
}

func TestFieldsAllocatesOnlyTheHeaders(t *testing.T) {
	line := strings.Repeat("column,", 63) + "last"
	if n := testing.AllocsPerRun(50, func() { _ = Fields(line, ',') }); n > 1 {
		t.Errorf("Fields made %v allocations, want 1: the pieces must be substrings", n)
	}
}
