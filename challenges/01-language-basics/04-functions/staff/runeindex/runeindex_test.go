package runeindex

import "testing"

func TestCharAt(t *testing.T) {
	if got := CharAt("héllo", 2); got != 'l' {
		t.Errorf("=%q want 'l' (byte indexing hits the wrong spot)", got)
	}
	if got := CharAt("日本語", 1); got != '本' {
		t.Errorf("=%q want '本'", got)
	}
}
