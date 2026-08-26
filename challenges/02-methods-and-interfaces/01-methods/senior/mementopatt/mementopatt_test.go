package mementopatt

import "testing"

func TestMemento(t *testing.T) {
	e := &Editor{Text: "initial"}
	m1 := e.Save()

	e.Text = "changed"
	e.Restore(m1)

	if e.Text != "initial" {
		t.Errorf("Text = %q, want initial", e.Text)
	}
}
