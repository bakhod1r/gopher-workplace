package errorsis

import (
	"errors"
	"testing"
)

func TestFetch(t *testing.T) {
	data := map[string]string{"k": "v"}

	if got, err := Fetch(data, "k"); got != "v" || err != nil {
		t.Errorf("Fetch = %q, %v", got, err)
	}

	_, err := Fetch(data, "missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want it to wrap ErrNotFound", err)
	}
	if got := err.Error(); got != "fetch missing: not found" {
		t.Errorf("Error() = %q, want \"fetch missing: not found\"", got)
	}
}

func TestIsMissing(t *testing.T) {
	_, err := Fetch(nil, "x")
	if !IsMissing(err) {
		t.Error("IsMissing on a wrapped ErrNotFound = false")
	}
	if IsMissing(nil) {
		t.Error("IsMissing(nil) = true")
	}
	if IsMissing(errors.New("other")) {
		t.Error("IsMissing on an unrelated error = true")
	}
}

func TestFetchAll(t *testing.T) {
	data := map[string]string{"a": "1", "b": "2"}

	got, err := FetchAll(data, []string{"a", "b"})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(got) != 2 || got[0] != "1" || got[1] != "2" {
		t.Errorf("FetchAll = %v", got)
	}

	if _, err := FetchAll(data, []string{"a", "zz"}); !IsMissing(err) {
		t.Errorf("err = %v, want a wrapped ErrNotFound", err)
	}

	got, err = FetchAll(data, nil)
	if err != nil || len(got) != 0 {
		t.Errorf("FetchAll(nil) = %v, %v", got, err)
	}
}
