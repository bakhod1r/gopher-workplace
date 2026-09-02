package slicescontainsfunc

import "testing"

func TestAnyExpired(t *testing.T) {
	if !AnyExpired([]Entry{{"a", 5}, {"b", 0}}) {
		t.Error("AnyExpired with a zero TTL = false, want true")
	}
	if !AnyExpired([]Entry{{"a", -1}}) {
		t.Error("AnyExpired with a negative TTL = false, want true")
	}
	if AnyExpired([]Entry{{"a", 5}, {"b", 1}}) {
		t.Error("AnyExpired with only live entries = true, want false")
	}
	if AnyExpired([]Entry{}) {
		t.Error("AnyExpired([]) = true, want false")
	}
}
