package cmporgen

import "testing"

func TestDisplay(t *testing.T) {
	if got := Display("nick", "user"); got != "nick" {
		t.Errorf(`Display("nick", "user") = %q, want "nick"`, got)
	}
	if got := Display("", "user"); got != "user" {
		t.Errorf(`Display("", "user") = %q, want "user"`, got)
	}
	if got := Display("", ""); got != "anonymous" {
		t.Errorf(`Display("", "") = %q, want "anonymous"`, got)
	}
	if got := Display("nick", ""); got != "nick" {
		t.Errorf(`Display("nick", "") = %q, want "nick"`, got)
	}
}
