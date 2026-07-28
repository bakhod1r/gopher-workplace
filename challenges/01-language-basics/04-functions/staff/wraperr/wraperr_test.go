package wraperr

import (
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	if err := Do(false); err != nil {
		t.Errorf("no failure: %v", err)
	}
	err := Do(true)
	if err == nil || !strings.HasPrefix(err.Error(), "do: ") {
		t.Errorf("want wrapped error, got %v", err)
	}
}
