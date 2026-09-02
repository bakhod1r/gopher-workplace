package statemachinegen

import "testing"

type state string
type event string

func TestMachineTransitions(t *testing.T) {
	m := NewMachine[state, event]("draft")
	m.Allow("draft", "submit", "sent")
	m.Allow("sent", "deliver", "done")

	if m.State() != "draft" {
		t.Fatalf("State() = %q, want draft", m.State())
	}
	if !m.Fire("submit") {
		t.Fatal(`Fire("submit") = false, want true`)
	}
	if m.State() != "sent" {
		t.Errorf("State() = %q, want sent", m.State())
	}
	if !m.Fire("deliver") {
		t.Error(`Fire("deliver") = false, want true`)
	}
	if m.State() != "done" {
		t.Errorf("State() = %q, want done", m.State())
	}
}

func TestMachineRejects(t *testing.T) {
	m := NewMachine[state, event]("draft")
	m.Allow("draft", "submit", "sent")
	if m.Fire("deliver") {
		t.Error(`Fire("deliver") from draft = true, want false`)
	}
	if m.State() != "draft" {
		t.Errorf("State() = %q, want draft (a rejected event must not move the machine)", m.State())
	}
	if m.Fire("unknown") {
		t.Error("an unknown event was accepted")
	}
}

func TestMachineIntStates(t *testing.T) {
	m := NewMachine[int, string](0)
	m.Allow(0, "go", 1)
	if !m.Fire("go") || m.State() != 1 {
		t.Errorf("State() = %d, want 1", m.State())
	}
}
