package hotpathshare

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestShare(t *testing.T) {
	flat := map[string]int64{"a": 3, "b": 1}
	if got := Share(flat, []string{"a"}); !near(got, 0.75) {
		t.Errorf("Share = %v, want 0.75", got)
	}
	if got := Share(flat, []string{"a", "b"}); !near(got, 1) {
		t.Errorf("Share = %v, want 1", got)
	}
	if got := Share(flat, nil); !near(got, 0) {
		t.Errorf("Share = %v, want 0", got)
	}
}

func TestShareIgnoresUnknownNames(t *testing.T) {
	flat := map[string]int64{"a": 3, "b": 1}
	if got := Share(flat, []string{"a", "nope"}); !near(got, 0.75) {
		t.Errorf("Share = %v, want 0.75", got)
	}
}

func TestShareCountsDuplicatesOnce(t *testing.T) {
	flat := map[string]int64{"a": 3, "b": 1}
	if got := Share(flat, []string{"a", "a", "a"}); !near(got, 0.75) {
		t.Errorf("Share = %v, want 0.75", got)
	}
}

func TestShareEmptyProfile(t *testing.T) {
	if got := Share(nil, []string{"a"}); got != 0 {
		t.Errorf("Share(nil, [a]) = %v, want 0", got)
	}
	if got := Share(map[string]int64{}, []string{"a"}); got != 0 {
		t.Errorf("Share(empty, [a]) = %v, want 0", got)
	}
}
